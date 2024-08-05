import { atom } from 'jotai';
import { atomWithQuery } from 'jotai-tanstack-query';
import axios from 'axios';

// Writable atoms
export const trendingSearchesAtom = atom([], {
  debugLabel: 'trendingSearchesAtom',
});

export const currentPageAtom = atom(0, {
  debugLabel: 'currentPageAtom',
});

export const fetchTrendingSearchesAtom = atomWithQuery(() => ({
  queryKey: ['trendingSearches'],
  queryFn: async () => {
    const response = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/trending`);
    return response.data;
  },
}), {
  debugLabel: 'fetchTrendingSearchesAtom',
});

export const userAtom = atom(null, {
  debugLabel: 'userAtom',
});
