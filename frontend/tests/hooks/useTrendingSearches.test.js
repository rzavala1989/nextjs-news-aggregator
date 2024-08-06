import { renderHook, waitFor } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import axios from 'axios';
import { useTrendingSearches } from '../../src/hooks/useTrendingSearches';

jest.mock('axios');

const queryClient = new QueryClient();

const mockData = [
  { query: 'Search 1' },
  { query: 'Search 2' },
  { query: 'Search 3' },
  { query: 'Search 4' },
  { query: 'Search 5' },
  { query: 'Search 6' },
  { query: 'Search 7' },
];

describe('useTrendingSearches Hook', () => {
  it('fetches and returns trending searches', async () => {
    axios.get.mockResolvedValue({ data: mockData });

    const { result } = renderHook(() => useTrendingSearches(), {
      wrapper: ({ children }) => <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>,
    });

    await waitFor(() => expect(result.current.data).toEqual(mockData));
    expect(result.current.isLoading).toBe(false);
  });
});
