/* eslint-disable react/no-unescaped-entities */
"use client";

import { useState } from 'react';
import { signIn } from 'next-auth/react';
import { useRouter } from 'next/navigation';
import useToast from '@/hooks/useToast';
import AuthFormLayout from '@/app/components/AuthFormLayout';
import Toast from '@/app/components/Toast';

export default function Login() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const router = useRouter();
  const { showToast } = useToast();

  const handleSubmit = async (e) => {
    e.preventDefault();
    const result = await signIn('credentials', {
      redirect: false,
      username,
      password,
    });

    if (result?.error) {
      showToast(result.error, 'error');
    } else {
      router.push('/news');
    }
  };

  return (
    <AuthFormLayout imageUrl="/images/login-image.jpg" formType="Login">
      <Toast />
      <form onSubmit={handleSubmit} className="flex flex-col space-y-4 mt-8">
        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          className="input input-bordered w-full"
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="input input-bordered w-full"
        />
        <button type="submit" className="btn btn-primary w-full">Login</button>
      </form>
    </AuthFormLayout>
  );
}
