import React, { useState } from "react";

const testimonials = [
  {
    name: "Jane Doe",
    role: "Software Engineer",
    review: "This platform helped me learn so much in a short time!",
    image: "https://via.placeholder.com/100", // Replace with actual image link
  },
  {
    name: "John Smith",
    role: "Product Designer",
    review: "Incredible experience. The design and structure are top-notch.",
    image: "https://via.placeholder.com/100",
  },
  {
    name: "Alice Johnson",
    role: "Data Analyst",
    review: "Highly recommend this to anyone looking to upskill.",
    image: "https://via.placeholder.com/100",
  },
];

const Testimonials = () => {
  const [currentIndex, setCurrentIndex] = useState(0);

  const prevTestimonial = () => {
    setCurrentIndex((prev) => (prev === 0 ? testimonials.length - 1 : prev - 1));
  };

  const nextTestimonial = () => {
    setCurrentIndex((prev) => (prev === testimonials.length - 1 ? 0 : prev + 1));
  };

  const { name, role, review, image } = testimonials[currentIndex];

  return (
    <section className="flex items-center justify-center py-10 bg-gray-100">
      <div className="relative bg-white p-8 rounded-lg shadow-md text-center max-w-md w-full">
        {/* Left arrow */}
        <button
          onClick={prevTestimonial}
          className="absolute left-[-2rem] top-1/2 transform -translate-y-1/2 text-3xl text-gray-500 hover:text-gray-700"
        >
          &#8592;
        </button>

        {/* Image */}
        <img
          src={image}
          alt={name}
          className="mx-auto rounded-full w-24 h-24 mb-4 border-4 border-blue-500"
        />

        {/* Name */}
        <h3 className="text-lg font-semibold">{name}</h3>
        {/* Role */}
        <p className="text-sm text-gray-500 mb-2">{role}</p>
        {/* Review */}
        <p className="text-gray-700 italic">"{review}"</p>

        {/* Right arrow */}
        <button
          onClick={nextTestimonial}
          className="absolute right-[-2rem] top-1/2 transform -translate-y-1/2 text-3xl text-gray-500 hover:text-gray-700"
        >
          &#8594;
        </button>
      </div>
    </section>
  );
};

export default Testimonials;
