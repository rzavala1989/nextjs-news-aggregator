// src/components/Navbar.jsx
"use client";

import Link from 'next/link';
import { useSession, signOut } from 'next-auth/react';
import useToast from '@/hooks/useToast';
import Toast from './Toast';

const Navbar = () => {
  const { data: session, status } = useSession();
  const { showToast } = useToast();

  const handleLogout = () => {
    signOut();
    showToast('Logged out successfully', 'success');
  };

  return (
    <>
      <Toast />
      <nav className="bg-gray-800 p-4">
        <div className="mx-auto flex justify-between items-center">
          <Link href="/" className="navbar-link text-xl">
            News Aggregator
          </Link>
          <div>
            {status === 'authenticated' ? (
              <>
                <span className="text-white mx-2">Welcome, {session.user.username}</span>
                <button onClick={handleLogout} className="navbar-link">Logout</button>
              </>
            ) : (
              <>
                <Link href="/login" className="navbar-link">Login</Link>
                <Link href="/register" className="navbar-link">Register</Link>
              </>
            )}
            <Link href="/news" className="navbar-link">News</Link>
          </div>
        </div>
      </nav>
    </>
  );
};

export default Navbar;
