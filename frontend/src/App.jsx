import React from "react";
import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Signup from "./pages/Signup";

function App() {
  return (
    <Router>
      <div className="min-h-screen flex flex-col">
        <Navbar />
        <main className="flex-grow">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/signup" element={<Signup />} />
          </Routes>
        </main>
        <Footer />
      </div>
    </Router>
  );
}

function Navbar() {
  return (
    <nav className="flex items-center justify-between p-4 bg-gray-800 text-white">
      <Link to="/" className="text-2xl font-bold hover:text-gray-300">
        LOGO
      </Link>
      <div>
        <Link to="/login" className="mr-4 hover:underline">
          Login
        </Link>
        <Link to="/signup" className="hover:underline">
          Sign Up
        </Link>
      </div>
    </nav>
  );
}

function Footer() {
  return (
    <footer className="p-4 bg-gray-900 text-white text-center">
      &copy; 2025 Your Website Name. All rights reserved.
    </footer>
  );
}

export default App;
