"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import { useAuth } from "@/context/AuthContext";

export default function SignInPage() {
  const router = useRouter();

  const { login } = useAuth();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const [error, setError] = useState("");

  const [loading, setLoading] = useState(false);


  const handleSubmit = async (
    e: React.FormEvent
  ) => {
    e.preventDefault();

    setError("");
    setLoading(true);

    try {

      await login({email, password});

      router.push("/");

    } catch (err) {

      setError(
        "Email atau password salah"
      );

    } finally {

      setLoading(false);

    }
  };


  return (
    <div className="flex min-h-screen items-center justify-center bg-gray-100 dark:bg-gray-900">

      <div className="w-full max-w-md rounded-xl bg-white p-8 shadow-lg dark:bg-gray-800">


        <h1 className="mb-6 text-2xl font-bold text-gray-900 dark:text-white">
          Sign In
        </h1>


        {
          error && (
            <div className="mb-4 rounded bg-red-100 p-3 text-red-600">
              {error}
            </div>
          )
        }



        <form
          onSubmit={handleSubmit}
          className="space-y-5"
        >


          <div>

            <label className="mb-2 block text-sm">
              Email
            </label>

            <input
              type="email"
              value={email}
              onChange={(e)=>setEmail(e.target.value)}
              className="w-full rounded border px-4 py-2"
              placeholder="email@example.com"
              required
            />

          </div>



          <div>

            <label className="mb-2 block text-sm">
              Password
            </label>


            <input
              type="password"
              value={password}
              onChange={(e)=>setPassword(e.target.value)}
              className="w-full rounded border px-4 py-2"
              placeholder="********"
              required
            />

          </div>

          <button
            type="submit"
            disabled={loading}
            className="w-full rounded bg-blue-600 py-2 text-white hover:bg-blue-700 disabled:bg-gray-400"
          >

            {
              loading
              ? "Loading..."
              : "Sign In"
            }
          </button>
        </form>

      </div>

    </div>
  );
}