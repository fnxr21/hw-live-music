"use client";

import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();

  return (
    <div className="flex min-h-screen items-center justify-center bg-zinc-50 dark:bg-black">
      <div className="w-full max-w-sm rounded-xl bg-white p-8 shadow-lg dark:bg-neutral-900">
        <h1 className="mb-2 text-center text-2xl font-bold">
          Live Music System
        </h1>
        <p className="mb-8 text-center text-sm text-neutral-500">
          Choose how you want to enter
        </p>

        <div className="space-y-4">
          {/* USER / TABLE LOGIN */}
          <button
            onClick={() => router.push("/table/2")}
            className="w-full rounded-lg bg-green-500 py-3 text-sm font-semibold text-black hover:bg-green-400"
          >
            Enter as Table User 🎧
          </button>

          {/* ADMIN LOGIN */}
          <button
            onClick={() => router.push("/dashboard")}
            className="w-full rounded-lg border border-neutral-300 py-3 text-sm font-semibold hover:bg-neutral-100 dark:border-neutral-700 dark:hover:bg-neutral-800"
          >
            Enter as Admin 🛠
          </button>
        </div>

        <p className="mt-6 text-center text-xs text-neutral-400">
          Demo mode · No authentication yet
        </p>
      </div>
    </div>
  );
}
