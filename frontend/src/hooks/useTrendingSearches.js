// src/hooks/useTrendingSearches.js
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';

export const useTrendingSearches = () => {
  return useQuery({
    queryKey: ['trendingSearches'],
    queryFn: async () => {
      const { data } = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/trending`);
      return data;
    },
    staleTime: 12 * 60 * 60 * 1000, // 12 hours
  });
};
