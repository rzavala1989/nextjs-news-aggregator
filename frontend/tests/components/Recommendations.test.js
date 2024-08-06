import '@testing-library/jest-dom';
import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import Recommendations from '../../src/app/components/Recommendations';
import { useRecommendations } from '../../src/hooks/useRecommendations';
import { useSession } from 'next-auth/react';

jest.mock('../../src/hooks/useRecommendations');
jest.mock('next-auth/react');

const mockRecommendations = [
  {
    id: 1,
    title: 'Test Article 1',
    content: 'Test Content 1',
    author: 'Test Author 1',
    url: 'http://test1.com',
    image_url: 'http://test1.com/image.jpg',
    published_at: '2024-08-04',
  },
  {
    id: 2,
    title: 'Test Article 2',
    content: 'Test Content 2',
    author: 'Test Author 2',
    url: 'http://test2.com',
    image_url: 'http://test2.com/image.jpg',
    published_at: '2024-08-04',
  },
];

describe('Recommendations', () => {
  let originalLog;

  beforeAll(() => {
    // Store the original console.log function
    originalLog = console.log;
    // Replace console.log with a jest mock function
    console.log = jest.fn();
  });

  afterAll(() => {
    // Restore the original console.log function
    console.log = originalLog;
  });

  beforeEach(() => {
    jest.clearAllMocks();
    useRecommendations.mockReturnValue({
      data: mockRecommendations,
      isLoading: false,
      error: null,
    });
    useSession.mockReturnValue({
      data: null,
      status: 'unauthenticated',
    });
  });

  it('renders Recommendations component', async () => {
    render(<Recommendations />);
    await waitFor(() => expect(screen.getByText('Recommendations')).toBeInTheDocument());
    expect(screen.getByText('Test Article 1')).toBeInTheDocument();
    expect(screen.getByText('Test Article 2')).toBeInTheDocument();
  });

  it('shows loading state initially', () => {
    useRecommendations.mockReturnValueOnce({
      data: null,
      isLoading: true,
      error: null,
    });
    render(<Recommendations />);
    expect(screen.getByText('Loading...')).toBeInTheDocument();
  });

  it('shows error state', () => {
    useRecommendations.mockReturnValueOnce({
      data: null,
      isLoading: false,
      error: 'Error loading recommendations',
    });
    render(<Recommendations />);
    expect(screen.getByText('Error loading recommendations')).toBeInTheDocument();
  });
});
