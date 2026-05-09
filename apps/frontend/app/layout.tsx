import "./globals.css";

import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "GameSense",
  description: "Gameplay intelligence for semantic VOD retrieval and AI coaching."
};

export default function RootLayout({
  children
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
