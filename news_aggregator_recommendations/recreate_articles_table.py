import os
from sqlalchemy import create_engine, Column, Integer, String, Text
from sqlalchemy.orm import declarative_base, sessionmaker, scoped_session
from dotenv import load_dotenv

load_dotenv()

Base = declarative_base()


class Article(Base):
    __tablename__ = 'articles'
    id = Column(Integer, primary_key=True)
    username = Column(String, nullable=False)
    title = Column(String, nullable=False)
    content = Column(Text, nullable=False)
    author = Column(String, nullable=True)
    url = Column(String, nullable=False)
    image_url = Column(String, nullable=True)
    language = Column(String, default='en')
    published_at = Column(String, nullable=False)


def setup_database():
    # Get the database path from the environment variable
    database_path = os.getenv('DATABASE_PATH', 'app.db')
    engine = create_engine(f'sqlite:///{database_path}')
    Base.metadata.create_all(engine)
    return scoped_session(sessionmaker(bind=engine))


if __name__ == "__main__":
    Session = setup_database()
    print("Articles table created successfully.")
