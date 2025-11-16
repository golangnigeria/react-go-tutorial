import { Routes, Route } from "react-router-dom";
import Navbar from "./components/layout/Navbar";
import Home from "./pages/home/Home";
 

export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:5000/api" : "/api";

export default function App() {
  return (
    <>
      <Navbar />

      <Routes>
        <Route path="/" element={<Home />} />
        
      </Routes>
    </>
  );
}
