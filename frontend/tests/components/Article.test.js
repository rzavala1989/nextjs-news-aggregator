import '@testing-library/jest-dom';
import React from 'react';
import { render, screen } from '@testing-library/react';
import Article from '../../src/app/components/Article';

const mockArticle = {
  title: 'Test Title',
  content: 'Test Content',
  author: 'Test Author',
  url: 'http://test.com',
  image_url: 'http://test.com/image.jpg',
  published_at: '2023-08-01T00:00:00Z'
};

describe('Article', () => {
  it('renders Article component with image', () => {
    render(<Article article={mockArticle} />);

    expect(screen.getByText(/Test Title/i)).toBeInTheDocument();
    expect(screen.getByText(/Test Content/i)).toBeInTheDocument();
    expect(screen.getByRole('link', { name: /Read more/i })).toHaveAttribute('href', 'http://test.com');
  });
});
