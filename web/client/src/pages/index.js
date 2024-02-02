import React from 'react';
import { Link, navigate } from 'gatsby';
import NavBar from "./components/Navbar"; 
import Footer from './components/Footer';

export default function Home() {
  return (
    <div className="font-sans text-gray-800 dark:text-gray-200">
      <NavBar />
      <section className="bg-white dark:bg-gray-900">
        <div className="container px-6 py-10 mx-auto">
          <h1 className="text-3xl font-semibold text-center text-gray-800 capitalize lg:text-4xl dark:text-white">
            explore our <br /> Collaborative Travel Diary <span className="underline decoration-blue-500">goTravel</span>
          </h1>

          <p className="max-w-2xl mx-auto mt-6 text-center text-gray-500 dark:text-gray-300">
            Discover the joy of shared memories with goTravel, the ultimate travel diary app designed for adventurers who love to explore the world together. Whether you're backpacking across continents or embarking on weekend getaways, goTravel brings your travel stories to life, allowing you and your friends to collaboratively document every step of your journey in a common diary.
          </p>

          <div class="container px-6 py-16 mx-auto text-center">
            <div class="max-w-lg mx-auto">
              <h4 class="text-3xl font-semibold text-gray-800 dark:text-white lg:text-2xl">Why?</h4>

              <p class="mt-6 text-gray-500 dark:text-gray-300">Capture breathtaking landscapes, jot down unforgettable experiences, and keep all your travel memories in one place. With goTravel, every trip becomes a shared adventure, creating a tapestry of memories that you and your friends will cherish forever. Start your journey with goTravel today, and turn your travels into tales worth telling.</p>
              <button onClick={() => navigate('/register')} className="px-5 py-2 mt-6 text-sm font-medium leading-5 text-center text-white capitalize bg-blue-600 rounded-lg hover:bg-blue-500 lg:mx-0 lg:w-auto focus:outline-none">
                Sign up here
              </button>

            </div>

            <div class="flex justify-center mt-10">
              <img class="object-cover w-full h-96 rounded-xl lg:w-4/5" src="https://images.unsplash.com/photo-1556761175-5973dc0f32e7?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1632&q=80" />
            </div>
          </div>

          <div className="grid grid-cols-1 gap-8 mt-8 md:grid-cols-2 lg:grid-cols-3">
            <div className="p-8 space-y-3 border-2 border-blue-500 dark:border-blue-300 rounded-xl">
              <span className="inline-block text-blue-500 dark:text-blue-400">
              </span>
              <h1 className="text-2xl font-semibold text-gray-700 capitalize dark:text-white">Collaborative</h1>
              <p className="text-gray-500 dark:text-gray-300">
                Going on trips with your friends? Add them and share a trip diary together!
              </p>
              <Link to="#" className="inline-flex items-center p-2 text-blue-500 capitalize transition-colors duration-300 transform bg-blue-100 rounded-md dark:bg-blue-500 dark:text-white hover:bg-blue-200 dark:hover:bg-blue-400">
                Learn More
              </Link>
            </div>

            <div className="p-8 space-y-3 border-2 border-blue-500 dark:border-blue-300 rounded-xl">
              <span className="inline-block text-blue-500 dark:text-blue-400">
                {/* SVG or Image Placeholder */}
              </span>
              <h1 className="text-2xl font-semibold text-gray-700 capitalize dark:text-white">Personal Diary</h1>
              <p className="text-gray-500 dark:text-gray-300">
                Going on solo trips? Create a personal diary where you can record your precious memories.
              </p>
              <Link to="#" className="inline-flex items-center p-2 text-blue-500 capitalize transition-colors duration-300 transform bg-blue-100 rounded-md dark:bg-blue-500 dark:text-white hover:bg-blue-200 dark:hover:bg-blue-400">
                Learn More
              </Link>
            </div>

            <div className="p-8 space-y-3 border-2 border-blue-500 dark:border-blue-300 rounded-xl">
              <span className="inline-block text-blue-500 dark:text-blue-400">
                {/* SVG or Image Placeholder */}
              </span>
              <h1 className="text-2xl font-semibold text-gray-700 capitalize dark:text-white">Social Media</h1>
              <p className="text-gray-500 dark:text-gray-300">
                Add your friends and see their travel trips, they can view yours too!
              </p>
              <Link to="#" className="inline-flex items-center p-2 text-blue-500 capitalize transition-colors duration-300 transform bg-blue-100 rounded-md dark:bg-blue-500 dark:text-white hover:bg-blue-200 dark:hover:bg-blue-400">
                Learn More
              </Link>
            </div>

          </div>
        </div>
      </section>

      <Footer />
    </div>
  );
}
