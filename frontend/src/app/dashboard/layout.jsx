"use client";

export default function AdminLayout({ children }) {
  return (
    <div className="min-h-screen bg-neutral-950 text-white flex">
      {children}
    </div>
  );
}
