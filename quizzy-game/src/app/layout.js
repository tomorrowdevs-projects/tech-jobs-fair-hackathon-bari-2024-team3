import { Inter } from "next/font/google";
import { ChakraProvider} from '@chakra-ui/react';
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Quizzy Games",
  description: "Hackaton TECH JOBS fair Bari 2024",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      
      <body className={inter.className}><ChakraProvider>{children}</ChakraProvider></body>
      
    </html>
  );
}
