# model.py
import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import linear_kernel
import pickle
from models.article import setup_database, Article

def update_similarity_matrix():
    session = setup_database()()

    # Load articles data
    articles_df = pd.read_sql(session.query(Article).statement, session.bind)

    # Prepare the TF-IDF matrix
    tfidf = TfidfVectorizer(stop_words='english')
    tfidf_matrix = tfidf.fit_transform(articles_df['content'])

    # Calculate the cosine similarity matrix
    cosine_sim = linear_kernel(tfidf_matrix, tfidf_matrix)

    # Save the model and data
    with open('model/tfidf_vectorizer.pkl', 'wb') as f:
        pickle.dump(tfidf, f)
    with open('model/cosine_similarity.pkl', 'wb') as f:
        pickle.dump(cosine_sim, f)
    articles_df.to_csv('model/articles.csv', index=False)

if __name__ == "__main__":
    update_similarity_matrix()
    print("Similarity matrix updated successfully.")
