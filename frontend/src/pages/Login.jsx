import React from "react";
import { Link } from "react-router-dom";

function Login() {
  return (
    <div className="p-8 flex items-center justify-between">
      <img src="/images/login.jpg" alt="Login" className="w-1/2 rounded-lg" />
      <div className="w-1/2 max-w-md mx-auto space-y-4">
        <h2 className="text-2xl font-bold">Login</h2>
        <input type="text" placeholder="Username" className="w-full p-2 border rounded" />
        <input type="password" placeholder="Password" className="w-full p-2 border rounded" />
        <button className="w-full bg-blue-600 text-white py-2 rounded">Login</button>
        <p className="text-center">
          Don’t have an account? <Link to="/signup" className="text-blue-600">Sign up</Link>
        </p>
      </div>
    </div>
  );
}

export default Login;
