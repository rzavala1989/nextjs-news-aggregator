import { useRecommendations } from '../../src/hooks/useRecommendations';
import { renderHook } from '@testing-library/react';

jest.mock('../../src/hooks/useRecommendations');

const mockRecommendations = [
  {
    id: 1,
    title: 'Test Title 1',
    content: 'Test Content 1',
    author: 'Author 1',
    url: 'http://test1.com',
    image_url: 'http://test1.com/image.jpg',
    published_at: '2024-08-05',
  },
  {
    id: 2,
    title: 'Test Title 2',
    content: 'Test Content 2',
    author: 'Author 2',
    url: 'http://test2.com',
    image_url: 'http://test2.com/image.jpg',
    published_at: '2024-08-06',
  },
];

describe('useRecommendations', () => {
  it('returns data', async () => {
    useRecommendations.mockReturnValue({
      data: mockRecommendations,
      isLoading: false,
      error: null,
    });

    const { result } = renderHook(() => useRecommendations(1, 10));

    expect(result.current.data).toEqual(mockRecommendations);
    expect(result.current.isLoading).toBeFalsy();
    expect(result.current.error).toBeNull();
  });

  it('shows loading state', async () => {
    useRecommendations.mockReturnValue({
      data: null,
      isLoading: true,
      error: null,
    });

    const { result } = renderHook(() => useRecommendations(1, 10));

    expect(result.current.data).toBeNull();
    expect(result.current.isLoading).toBeTruthy();
    expect(result.current.error).toBeNull();
  });

  it('handles error state', async () => {
    useRecommendations.mockReturnValue({
      data: null,
      isLoading: false,
      error: 'Error fetching recommendations',
    });

    const { result } = renderHook(() => useRecommendations(1, 10));

    expect(result.current.data).toBeNull();
    expect(result.current.isLoading).toBeFalsy();
    expect(result.current.error).toEqual('Error fetching recommendations');
  });
});
