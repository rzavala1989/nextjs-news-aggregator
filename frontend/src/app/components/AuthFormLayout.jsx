/* eslint-disable react/no-unescaped-entities */
"use client";

import Image from 'next/image';
import Link from 'next/link';

const AuthFormLayout = ({ children, imageUrl, formType }) => {
  return (
    <div className="min-h-screen flex">
      <div className="w-2/3 relative">
        <Image
          src={imageUrl}
          alt={`${formType} image`}
          layout="fill"
          objectFit="cover"
        />
        <div className="absolute inset-0 bg-black opacity-50"></div>
        <div className="absolute inset-0 flex items-center justify-center text-white text-center">
          <h1 className="text-5xl font-bold">Hello World.</h1>
        </div>
      </div>
      <div className="w-1/3 flex items-center justify-center p-8 bg-gray-100">
        <div className="w-full max-w-md">
          <h2 className="text-3xl font-bold mb-4">{formType}</h2>
          {children}
          <div className="mt-4">
            {formType === "Login" ? (
              <p>
                Don't have an account?{' '}
                <Link href="/register" className="text-blue-500">
                  Create your account
                </Link>
              </p>
            ) : (
              <p>
                Already have an account?{' '}
                <Link href="/login" className="text-blue-500">
                  Login
                </Link>
              </p>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default AuthFormLayout;
