# models/article.py
import os
from sqlalchemy import create_engine, Column, Integer, String, Text
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker, scoped_session
from dotenv import load_dotenv

Base = declarative_base()

load_dotenv()

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
    database_path = os.getenv('DATABASE_PATH')
    engine = create_engine(f'sqlite:///{database_path}')
    Base.metadata.create_all(engine)
    session_factory = sessionmaker(bind=engine)
    return scoped_session(session_factory)
