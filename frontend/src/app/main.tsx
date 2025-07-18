/**
 * Main Application Entry Point
 * 
 * This file serves as the primary entry point for the Minden Airport Frontend application.
 * It sets up the React Router configuration and renders the main application routes.
 * 
 * The application supports both web and Electron modes, providing:
 * - Public home page with flight information
 * - User authentication (login/register)
 * - User dashboard for ticket and baggage management
 * - Admin interface for flight and system management
 * 
 * @author Giulien Chow
 * @version 1.0.0
 */

import ReactDOM from "react-dom/client";
import "@/assets/styles.css";
import { BrowserRouter, Route, Routes } from "react-router";
import Home from "@/features/home";
import { Login, Register } from "@/features/auth";
import Dashboard from "@/features/dashboard";
import { AdminPage } from "@/features/admin";

// Get the root DOM element where the React app will be mounted
const root = document.getElementById("root");

if (root) {
  // Initialize React application with routing configuration
  ReactDOM.createRoot(root).render(
    <BrowserRouter>
      <Routes>
        <Route>
          {/* Public home page - displays flight information and airport status */}
          <Route path="/" element={<Home />} />
          
          {/* Authentication routes */}
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          
          {/* Protected user dashboard - requires authentication */}
          <Route path="/dashboard" element={<Dashboard />} />
          
          {/* Admin interface - requires admin role */}
          <Route path="/admin" element={<AdminPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
} else {
  throw new Error("Root element not found");
}
