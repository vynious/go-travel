// Import React and any other necessary libraries
import * as React from 'react';
import { Link } from 'gatsby'; // Import Link component for navigation
import NavBar from './components/Navbar';

// Define the Home component
export default function Home() {
  return (
    <div style={{ fontFamily: 'sans-serif', textAlign: 'center', padding: '50px' }}>
      {/* Application Header */}
      


      {/* Main Content Section */}
      <main>
        <NavBar />

        <section>
          <h2>Discover the World Through the Eyes of Fellow Travelers</h2>
          <p>Join our community of explorers sharing their unique travel experiences from around the globe.</p>
        </section>

        <section>
          <h2>Create Your Own Travel Diary</h2>
          <p>Document your adventures, share your journey, and connect with like-minded individuals.</p>
        </section>

        {/* Call to Action Buttons */}
        <div style={{ marginTop: '30px' }}>
          <Link to="/diaries" style={{ marginRight: '20px', textDecoration: 'none', color: 'white', background: 'blue', padding: '10px 20px', borderRadius: '5px' }}>Explore Diaries</Link>
          <Link to="/join" style={{ textDecoration: 'none', color: 'white', background: 'green', padding: '10px 20px', borderRadius: '5px' }}>Join Now</Link>
        </div>
      </main>

      {/* Footer */}
      <footer style={{ marginTop: '50px' }}>
        <p>© {new Date().getFullYear()} Wanderlust Diaries. All rights reserved.</p>
      </footer>
    </div>
  );
}
