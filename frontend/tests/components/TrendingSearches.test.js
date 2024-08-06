import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import TrendingSearches from '../../src/app/components/TrendingSearches';
import { useTrendingSearches } from '../../src/hooks/useTrendingSearches';

jest.mock('../../src/hooks/useTrendingSearches');

const mockData = [
  { query: 'Search 1' },
  { query: 'Search 2' },
  { query: 'Search 3' },
  { query: 'Search 4' },
  { query: 'Search 5' },
  { query: 'Search 6' },
  { query: 'Search 7' },
];

describe('TrendingSearches Component', () => {
  beforeEach(() => {
    useTrendingSearches.mockReturnValue({ data: mockData, isLoading: false, error: null });
  });

  it('renders TrendingSearches component', () => {
    render(<TrendingSearches />);

    expect(screen.getByText('Trending Searches')).toBeInTheDocument();
    expect(screen.getByText('Search 1')).toBeInTheDocument();
    expect(screen.getByText('Search 2')).toBeInTheDocument();
  });

  it('pagination works correctly', async () => {
    render(<TrendingSearches />);

    const nextButton = screen.getByText('Next');
    fireEvent.click(nextButton);

    await waitFor(() => expect(screen.getByText('Search 6')).toBeInTheDocument());

    expect(screen.queryByText('Search 1')).not.toBeInTheDocument();
  });

  it('shows loading state initially', () => {
    useTrendingSearches.mockReturnValueOnce({
      data: null,
      isLoading: true,
      error: null,
    });
    render(<TrendingSearches />);
    expect(screen.getByText('Loading...')).toBeInTheDocument();
  });
});
