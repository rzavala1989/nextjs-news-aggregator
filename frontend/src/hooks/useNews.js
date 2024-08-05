// src/hooks/useNews.js
import { useQuery, useMutation } from '@tanstack/react-query';
import axios from 'axios';
import { useSession } from 'next-auth/react';

const fetchNews = async ({ queryKey }) => {
  const [_, token, page, limit] = queryKey;
  const { data } = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/secure/news`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
    params: {
      page,
      limit,
    },
  });
  return data;
};

const syncArticles = async (token) => {
  const { data } = await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/secure/sync`, {}, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return data;
};

export const useNews = (page, limit) => {
  const { data: session } = useSession();
  const token = session?.user?.token;

  return useQuery({
    queryKey: ['news', token, page, limit],
    queryFn: fetchNews,
    enabled: !!token, // Only run the query if there's a token
  });
};

export const useSyncArticles = () => {
  const { data: session } = useSession();
  const token = session?.user?.token;

  return useMutation(() => syncArticles(token));
};
