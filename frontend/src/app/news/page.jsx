// src/app/news/page.jsx
"use client";

import { useRecommendations } from '@/hooks/useRecommendations';
import { useSession, signIn } from 'next-auth/react';
import { useEffect, useState } from 'react';
import useToast from '@/hooks/useToast';
import Toast from '@/app/components/Toast';
import Article from '@/app/components/Article';
import 'daisyui';

export default function News() {
  const { data: session, status } = useSession();
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 10;
  const { data: recommendations, isLoading, error } = useRecommendations(currentPage, itemsPerPage);
  const { showToast } = useToast();

  useEffect(() => {
    console.log('Session:', session);
    if (status === 'unauthenticated') {
      showToast('Please sign in to view the news', 'info');
      signIn();
    }
  }, [status, showToast, session]);

  const handleNextPage = () => {
    setCurrentPage((prevPage) => prevPage + 1);
  };

  const handlePrevPage = () => {
    setCurrentPage((prevPage) => Math.max(prevPage - 1, 1));
  };

  if (isLoading) return <p>Loading...</p>;
  if (error) return <p>Error loading news</p>;

  return (
    <>
      <Toast />
      <div className="container mx-auto py-4">
        <h1 className="text-2xl font-bold mb-4">News</h1>
        {recommendations?.length > 0 ? (
          <>
            <ul className="space-y-4">
              {recommendations.map((article) => (
                <Article key={article.id} article={article} />
              ))}
            </ul>
            <div className="flex justify-between mt-4">
              <button
                onClick={handlePrevPage}
                disabled={currentPage === 1}
                className="btn btn-primary"
              >
                Previous
              </button>
              <button
                onClick={handleNextPage}
                disabled={recommendations.length < itemsPerPage}
                className="btn btn-primary"
              >
                Next
              </button>
            </div>
          </>
        ) : (
          <p>No articles available</p>
        )}
      </div>
    </>
  );
}
