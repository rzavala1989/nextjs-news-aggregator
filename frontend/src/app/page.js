// src/app/page.js
"use client";

import { useEffect } from 'react';
import { useTrendingSearches } from '@/hooks/useTrendingSearches';
import useToast from '../hooks/useToast';
import Toast from '@/app/components/Toast';
import Recommendations from '@/app/components/Recommendations';

const HomePage = () => {
  const { data: trendingSearches, isLoading, error } = useTrendingSearches();
  const { showToast } = useToast();

  useEffect(() => {
    if (error) {
      showToast('Failed to fetch trending searches', 'error');
    }
  }, [error, showToast]);

  return (
    <>
      <Toast />
      <div className="container mx-auto py-4">
        <h1 className="text-2xl font-bold mb-4 text-center">Welcome to the News Aggregator</h1>
        {isLoading ? (
          <p>Loading...</p>
        ) : (
          <>
            {trendingSearches && (
              <div className="mt-8">
                <h2 className="text-xl font-bold">Trending Searches</h2>
                <ul>
                  {trendingSearches.map((search, index) => (
                    <li key={index} className="text-red-500 whitespace-nowrap">{search.query}</li>
                  ))}
                </ul>
                <h2 className="text-xl font-bold mt-8">Recommendations</h2>
                <Recommendations />
              </div>
            )}
          </>
        )}
      </div>
    </>
  );
};

export default HomePage;
