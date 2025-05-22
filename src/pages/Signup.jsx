import React from "react";
import { Link } from "react-router-dom";

function Signup() {
  return (
    <div className="p-8 flex items-center justify-between">
      <img src="/images/signup.jpg" alt="Signup" className="w-1/2 rounded-lg" />
      <div className="w-1/2 max-w-md mx-auto space-y-4">
        <h2 className="text-2xl font-bold">Sign Up</h2>
        <input type="text" placeholder="Username" className="w-full p-2 border rounded" />
        <input type="email" placeholder="Email" className="w-full p-2 border rounded" />
        <input type="password" placeholder="Password" className="w-full p-2 border rounded" />
        <input type="password" placeholder="Confirm Password" className="w-full p-2 border rounded" />
        <button className="w-full bg-green-600 text-white py-2 rounded">Sign Up</button>
        <p className="text-sm text-center text-blue-600 cursor-pointer">Forgot password?</p>
        <p className="text-center">
          Already have an account? <Link to="/login" className="text-blue-600">Log in</Link>
        </p>
      </div>
    </div>
  );
}

export default Signup;
