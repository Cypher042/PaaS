"use client";

import { useAuth } from "@/hooks/useAuth";

export function LoginPage() {
  const { loginWithGithub } = useAuth();

  return (
    <div className="flex flex-col space-y-4 p-8 bg-white border rounded shadow-sm w-96 text-center">
      <h1 className="text-2xl font-bold">Welcome Back</h1>
      <p className="text-sm text-gray-500">Log in to your PaaS account</p>
      
      <button 
        onClick={loginWithGithub}
        className="w-full px-4 py-2 mt-4 text-white bg-gray-900 rounded hover:bg-gray-800 transition-colors"
      >
        Sign in with GitHub
      </button>
    </div>
  );
}
