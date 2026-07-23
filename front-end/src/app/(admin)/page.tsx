import type { Metadata } from "next";
import React from "react";

export const metadata: Metadata = {
  title: "Dashboard | YouTube Shorts Automation",
  description: "Dashboard YouTube Shorts Automation System",
};

export default function DashboardPage() {
  return (
    <div className="space-y-6">
      {/* Header */}
      <div>
        <h1 className="text-3xl font-bold text-gray-900 dark:text-white">
          Dashboard
        </h1>
        <p className="mt-2 text-gray-500 dark:text-gray-400">
          Welcome to the YouTube Shorts Automation System.
        </p>
      </div>

      {/* Statistic Cards */}
      <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 xl:grid-cols-4">
        <div className="rounded-xl border border-gray-200 bg-white p-6 shadow-sm dark:border-gray-800 dark:bg-gray-900">
          <h2 className="text-sm font-medium text-gray-500">Total Users</h2>
          <p className="mt-3 text-3xl font-bold">0</p>
        </div>

        <div className="rounded-xl border border-gray-200 bg-white p-6 shadow-sm dark:border-gray-800 dark:bg-gray-900">
          <h2 className="text-sm font-medium text-gray-500">Total Jobs</h2>
          <p className="mt-3 text-3xl font-bold">0</p>
        </div>

        <div className="rounded-xl border border-gray-200 bg-white p-6 shadow-sm dark:border-gray-800 dark:bg-gray-900">
          <h2 className="text-sm font-medium text-gray-500">Running Jobs</h2>
          <p className="mt-3 text-3xl font-bold">0</p>
        </div>

        <div className="rounded-xl border border-gray-200 bg-white p-6 shadow-sm dark:border-gray-800 dark:bg-gray-900">
          <h2 className="text-sm font-medium text-gray-500">History Records</h2>
          <p className="mt-3 text-3xl font-bold">0</p>
        </div>
      </div>

      {/* Recent Jobs */}
      <div className="rounded-xl border border-gray-200 bg-white p-6 shadow-sm dark:border-gray-800 dark:bg-gray-900">
        <h2 className="mb-4 text-xl font-semibold">Recent Jobs</h2>

        <div className="overflow-x-auto">
          <table className="min-w-full border-collapse">
            <thead>
              <tr className="border-b">
                <th className="px-4 py-3 text-left">Job Name</th>
                <th className="px-4 py-3 text-left">Status</th>
                <th className="px-4 py-3 text-left">Created At</th>
              </tr>
            </thead>

            <tbody>
              <tr>
                <td
                  colSpan={3}
                  className="px-4 py-8 text-center text-gray-500"
                >
                  No jobs available.
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}