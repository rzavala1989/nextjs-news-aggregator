/* eslint-disable react/no-unescaped-entities */
"use client";

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import axios from 'axios';
import AuthFormLayout from '@/app/components/AuthFormLayout';

export default function Register() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const router = useRouter();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/register`, { username, password });
      router.push('/login');
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <AuthFormLayout imageUrl="/images/register-image.jpg" formType="Register">
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
        <button type="submit" className="btn btn-primary w-full">
          Register
        </button>
      </form>
    </AuthFormLayout>
  );
}
