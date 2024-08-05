// src/components/Layout.jsx
"use client";

import { useEffect } from "react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { usePathname, useRouter } from 'next/navigation';
import { SessionProvider } from "next-auth/react";
import Navbar from './Navbar';
import TrendingSearches from './TrendingSearches';

const queryClient = new QueryClient();

const Layout = ({ children, session }) => {
  const router = useRouter();
  const pathname = usePathname();
  const isAuthPage = pathname === '/login' || pathname === '/register';

  useEffect(() => {
    console.log('Router initialized in Layout:', router);
  }, [router]);

  return (
    <SessionProvider session={session}>
      <QueryClientProvider client={queryClient}>
        <div>
          <Navbar />
          <div className="h-1 bg-accent my-0" />
          {!isAuthPage && <TrendingSearches />}
          <main className={isAuthPage ? "min-h-screen" : "container mx-auto py-4"}>
            {children}
          </main>
        </div>
        <ReactQueryDevtools initialIsOpen={false} />
      </QueryClientProvider>
    </SessionProvider>
  );
};

export default Layout;
