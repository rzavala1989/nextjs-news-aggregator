import create from 'zustand';
import { persist } from 'zustand/middleware';

const useNewsStore = create(
  persist(
    (set) => ({
      trendingSearches: [],
      currentPage: 0,
      setTrendingSearches: (searches) => set({ trendingSearches: searches }),
      setCurrentPage: (page) => set({ currentPage: page }),
    }),
    {
      name: 'news-storage',
    }
  )
);

export default useNewsStore;
