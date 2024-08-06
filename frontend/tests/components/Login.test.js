import '@testing-library/jest-dom';
import React from 'react';
import { render, screen } from '@testing-library/react';
import Login from '../../src/app/login/page';

jest.mock('@/hooks/useToast', () => ({
  __esModule: true,
  default: () => ({
    showToast: jest.fn(),
  }),
}));

jest.mock('next/navigation', () => ({
  useRouter: jest.fn(),
}));

describe('Login', () => {
  it('renders Login component', () => {
    render(<Login />);
    expect(screen.getByText('Login')).toBeInTheDocument();
  });
});
