// src/app/pages/api/auth/[...nextauth].js

import NextAuth from 'next-auth';
import CredentialsProvider from 'next-auth/providers/credentials';
import axios from 'axios';

export default NextAuth({
  providers: [
    CredentialsProvider({
      name: 'Credentials',
      credentials: {
        username: { label: 'Username', type: 'text' },
        password: { label: 'Password', type: 'password' },
      },
      async authorize(credentials) {
        try {
          const res = await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/login`, {
            username: credentials.username,
            password: credentials.password,
          });

          const { user, token } = res.data;

          if (user && token) {
            return { ...user, token }; // Return the user and the token
          } else {
            return null;
          }
        } catch (error) {
          throw new Error('Login failed');
        }
      },
    }),
  ],
  secret: process.env.NEXT_JWT_SECRET,
  callbacks: {
    async jwt({ token, user }) {
      if (user) {
        token.id = user.id;
        token.username = user.username;
        token.token = user.token; // Store the token in the JWT
      }
      return token;
    },
    async session({ session, token }) {
      session.user.id = token.id;
      session.user.username = token.username;
      session.user.token = token.token; // Include the token in the session
      return session;
    },
  },
  pages: {
    signIn: '/login',
    error: '/login',
  },
});
