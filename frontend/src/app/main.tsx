import ReactDOM from "react-dom/client";
import "@/assets/styles.css";
import { BrowserRouter, Route, Routes } from "react-router";
import Home from "@/features/home";
import { Login, Register } from "@/features/auth";
import Dashboard from "@/features/dashboard";
import { AdminPage } from "@/features/admin";

const root = document.getElementById("root");

if (root) {
  ReactDOM.createRoot(root).render(
    <BrowserRouter>
      <Routes>
        <Route>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/admin" element={<AdminPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
} else {
  throw new Error("Root element not found");
}
