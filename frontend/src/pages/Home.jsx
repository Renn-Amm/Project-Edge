import React from "react";

function Home() {
  return (
    <div className="p-8 space-y-12">
      <section className="flex items-center space-x-6">
        <div className="w-1/2">
          <h1 className="text-4xl font-bold mb-4">Hero Section</h1>
          <p>Welcome to our platform! This is the hero description.</p>
        </div>
        <img src="/images/hero.jpg" alt="Hero" className="w-1/2 rounded-lg" />
      </section>

      <section>
        <h2 className="text-3xl font-semibold mb-2">Our Vision and Mission</h2>
        <p>Our vision is to ... Our mission is to ...</p>
      </section>

      <section className="flex items-center space-x-6">
        <div className="w-1/2">
          <h2 className="text-3xl font-semibold mb-2">About Us</h2>
          <p>We are a team that ...</p>
        </div>
        <img src="/images/about.jpg" alt="About" className="w-1/2 rounded-lg" />
      </section>

      <section>
        <h2 className="text-3xl font-semibold mb-2">Testimonials</h2>
        <p>"This is the best platform ever!" - User A</p>
        <p>"Really helped me learn!" - User B</p>
      </section>
    </div>
  );
}

export default Home;
