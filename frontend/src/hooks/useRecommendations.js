// src/hooks/useRecommendations.js
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { useSession } from 'next-auth/react';

const fetchRecommendations = async ({ queryKey }) => {
  const [_, token, page, limit] = queryKey;
  const { data } = await axios.get(`${process.env.NEXT_PUBLIC_API_URL}/secure/recommendations`, {
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

export const useRecommendations = (page, limit) => {
  const { data: session } = useSession();
  const token = session?.user?.token;

  return useQuery({
    queryKey: ['recommendations', token, page, limit],
    queryFn: fetchRecommendations,
    enabled: !!token, // Only run the query if there's a token
  });
};
