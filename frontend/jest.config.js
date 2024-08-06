module.exports = {
  setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
  moduleNameMapper: {
    '^@/hooks/(.*)$': '<rootDir>/src/hooks/$1',
    '^@/app/components/(.*)$': '<rootDir>/src/app/components/$1',
    '^@/app/hooks/(.*)$': '<rootDir>/src/app/hooks/$1',
    '^@/app/pages/(.*)$': '<rootDir>/src/app/pages/$1',
    '^@/app/styles/(.*)$': '<rootDir>/src/app/styles/$1',
    '^@/app/utils/(.*)$': '<rootDir>/src/app/utils/$1',
    '^@/app/layouts/(.*)$': '<rootDir>/src/app/layouts/$1',
    '^next/router$': '<rootDir>/__mocks__/next/router.js',
    '\\.(css|less|scss|sass)$': 'jest-transform-stub',
  },
  transform: {
    '^.+\\.(js|jsx|ts|tsx)$': 'babel-jest',
  },
  testEnvironment: 'jsdom',
};
