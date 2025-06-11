import ReactDOM from "react-dom/client";
import "@/assets/styles.css";
import { BrowserRouter, Route, Routes } from "react-router";
import Home from "@/features/home";
import Layout from "@/components/layout";
import { Login, Register } from "@/features/auth";

const root = document.getElementById("root");

if (root) {
  ReactDOM.createRoot(root).render(
    <BrowserRouter>
      <Routes>
        <Route element={<Layout />}>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
} else {
  throw new Error("Root element not found");
}
