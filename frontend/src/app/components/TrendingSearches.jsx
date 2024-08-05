"use client";

import { useState, useEffect } from 'react';
import { useTrendingSearches } from '@/hooks/useTrendingSearches';

const TrendingSearches = ({ className }) => {
  const { data, isLoading, error } = useTrendingSearches();
  const [currentPage, setCurrentPage] = useState(0);
  const [transition, setTransition] = useState('');

  const itemsPerPage = 5;

  useEffect(() => {
    if (data) {
      // Data handling logic, if any additional is required
    }
  }, [data]);

  const paginatedSearches = data ? data.slice(
    currentPage * itemsPerPage,
    (currentPage + 1) * itemsPerPage
  ) : [];

  const handleNext = () => {
    if ((currentPage + 1) * itemsPerPage < data.length) {
      setTransition('slide-next');
      setTimeout(() => {
        setCurrentPage(currentPage + 1);
        setTransition('');
      }, 500); // Duration of the slide animation
    }
  };

  const handlePrev = () => {
    if (currentPage > 0) {
      setTransition('slide-prev');
      setTimeout(() => {
        setCurrentPage(currentPage - 1);
        setTransition('');
      }, 500); // Duration of the slide animation
    }
  };

  if (isLoading) return <p>Loading...</p>;
  if (error) return <p>Error loading trending searches</p>;

  return (
    <div className={`bg-darkBackground p-4 ${className}`}>
      <div className="flex justify-between items-center mb-4">
        <button
          onClick={handlePrev}
          disabled={currentPage === 0}
          className={`px-4 py-2 bg-buttonBg text-white rounded border-2 border-amber-200 drop-shadow-lg hover:bg-buttonHover transition-colors duration-300 ${currentPage === 0 && 'opacity-50 cursor-not-allowed'}`}
        >
          Prev
        </button>
        <div className="flex flex-grow justify-between mx-5 items-center space-x-4">
          <div className="flex flex-col items-center text-gray-100">
            <span className="text-lg font-bold">Trending Searches</span>
            <span className="text-sm">(as of {new Date().toLocaleDateString()})</span>
          </div>
          <div className={`flex space-x-4 overflow-hidden ${transition}`}>
            {paginatedSearches.map((search, index) => (
              <span key={index} className="text-white whitespace-nowrap bg-darkBackground py-2 px-4 rounded-full border border-white">
                {search.query}
              </span>
            ))}
          </div>
        </div>
        <button
          onClick={handleNext}
          disabled={(currentPage + 1) * itemsPerPage >= data.length}
          className={`px-4 py-2 bg-buttonBg text-white rounded border-2 border-amber-200 drop-shadow-lg hover:bg-buttonHover transition-colors duration-300 ${(currentPage + 1) * itemsPerPage >= data.length && 'opacity-50 cursor-not-allowed'}`}
        >
          Next
        </button>
      </div>
    </div>
  );
};

export default TrendingSearches;
