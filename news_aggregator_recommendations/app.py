import os
import pandas as pd
import pickle
from flask import Flask, request, jsonify
from sqlalchemy.orm import scoped_session

from model import update_similarity_matrix
from models.article import setup_database, Article

app = Flask(__name__)
Session = setup_database()

# Load the trained model and articles data
with open('model/tfidf_vectorizer.pkl', 'rb') as f:
    tfidf = pickle.load(f)
with open('model/cosine_similarity.pkl', 'rb') as f:
    cosine_sim = pickle.load(f)

@app.route('/recommend', methods=['POST'])
def recommend():
    session = Session()
    title = request.json.get('title')
    article = session.query(Article).filter_by(title=title).first()
    if not article:
        session.close()
        app.logger.error(f"Article with title '{title}' not found")
        return jsonify({'message': 'Article not found'}), 404

    try:
        articles_df = pd.read_sql(session.query(Article).statement, session.bind)
        idx = articles_df[articles_df['title'] == title].index[0]
        sim_scores = list(enumerate(cosine_sim[idx]))
        sim_scores = sorted(sim_scores, key=lambda x: x[1], reverse=True)
        sim_scores = sim_scores[1:6]  # Get top 5 similar articles
        article_indices = [i[0] for i in sim_scores]
        recommended_articles = articles_df.iloc[article_indices].to_dict('records')
        app.logger.info(f"Recommendations for '{title}': {[a['title'] for a in recommended_articles]}")
        return jsonify(recommended_articles)
    except IndexError:
        app.logger.error(f"Article '{title}' not found in similarity matrix")
        return jsonify({'message': 'Article not found in similarity matrix'}), 404
    finally:
        session.close()

@app.route('/sync', methods=['POST'])
def sync_articles():
    session = Session()
    articles_data = request.json
    new_articles_count = 0
    for article_data in articles_data:
        article = session.query(Article).filter_by(url=article_data['url']).first()
        if article:
            continue
        new_article = Article(
            username=article_data.get('username', 'testuser'),
            title=article_data['title'],
            content=article_data['content'],
            author=article_data.get('author', 'N/A'),
            url=article_data['url'],
            image_url=article_data.get('image_url', 'N/A'),
            language=article_data['language'],
            published_at=article_data['published_at']
        )
        session.add(new_article)
        new_articles_count += 1
    session.commit()
    session.close()

    if new_articles_count > 0:
        update_similarity_matrix()

    app.logger.info(f"Articles synced successfully. New articles added: {new_articles_count}")
    return jsonify({'message': f'Articles synced successfully. New articles added: {new_articles_count}'}), 200

if __name__ == '__main__':
    app.run(debug=True)
