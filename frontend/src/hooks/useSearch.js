// src/hooks/useSearch.js
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { useSession } from 'next-auth/react';

const fetchSearchResults = async ({ queryKey }) => {
  const [_, token, query] = queryKey;
  const { data } = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/search`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
    params: {
      query,
    },
  });
  return data;
};

export const useSearch = (query) => {
  const { data: session } = useSession();
  const token = session?.user?.token;

  return useQuery({
    queryKey: ['search', token, query],
    queryFn: fetchSearchResults,
    enabled: !!token && !!query, // Only run the query if there's a token and a search query
  });
};
