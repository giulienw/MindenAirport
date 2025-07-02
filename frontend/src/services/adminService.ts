import type {
  AdminDashboard,
  FlightManagement,
  BaggageManagement,
  User,
} from "@/types";
import { API_BASE_URL } from "@/config";
import { flightService } from "./flightService";

export const adminService = {
  async getAdminDashboard(): Promise<AdminDashboard> {
    const token = localStorage.getItem("token");
    if (!token) {
      throw new Error("No authentication token found");
    }

    try {
      const response = await fetch(`${API_BASE_URL}/admin/dashboard`, {
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to fetch admin dashboard");
      }

      const [dashboardResponse, flightsResponse] = await Promise.all([
        response.json(),
        flightService.getEnrichedFlights(),
      ]);
      return {
        stats: dashboardResponse.data.stats,
        airlines: dashboardResponse.data.airlines,
        airports: dashboardResponse.data.airports,
        flights: flightsResponse,
      } as AdminDashboard;
    } catch (error) {
      console.error("Failed to fetch admin dashboard:", error);
      throw error;
    }
  },

  async getUsers(
    page = 1,
    limit = 20,
    search?: string
  ): Promise<{ users: User[]; total: number }> {
    const token = localStorage.getItem("token");
    if (!token) {
      throw new Error("No authentication token found");
    }

    try {
      const params = new URLSearchParams({
        page: page.toString(),
        limit: limit.toString(),
        ...(search && { search }),
      });

      const response = await fetch(`${API_BASE_URL}/admin/users?${params}`, {
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to fetch users");
      }

      const result = await response.json();
      return result.data;
    } catch (error) {
      console.error("Failed to fetch users:", error);
      throw error;
    }
  },

  async updateUserStatus(userId: string, active: boolean): Promise<void> {
    const token = localStorage.getItem("token");
    if (!token) {
      throw new Error("No authentication token found");
    }

    try {
      const activeInt = active ? 1 : 0;
      const response = await fetch(`${API_BASE_URL}/admin/users/${userId}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ active: activeInt }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to update user status");
      }
    } catch (error) {
      console.error("Failed to update user status:", error);
      throw error;
    }
  },

  async getFlightManagement(): Promise<FlightManagement[]> {
    const token = localStorage.getItem("token");
    if (!token) {
      throw new Error("No authentication token found");
    }

    try {
      const response = await fetch(`${API_BASE_URL}/flight`, {
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(
          errorData.error || "Failed to fetch flight management data"
        );
      }

      const result = await response.json();
      return result;
    } catch (error) {
      console.error("Failed to fetch flight management data:", error);
      throw error;
    }
  },

  async updateFlightStatus(flightId: string, status: number): Promise<void> {
    const token = localStorage.getItem("token");
    if (!token) {
      throw new Error("No authentication token found");
    }

    try {
      const flight = await flightService.getFlightById(flightId);
      if (flight) {
        flight.statusId = status;
      }
      const response = await fetch(
        `${API_BASE_URL}/admin/flights/${flightId}`,
        {
          method: "PATCH",
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
          body: JSON.stringify(flight),
        }
      );

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to update flight status");
      }
    } catch (error) {
      console.error("Failed to update flight status:", error);
      throw error;
    }
  },

  async getBaggageManagement(): Promise<BaggageManagement[]> {
    const token = localStorage.getItem("token");
    if (!token) {
      throw new Error("No authentication token found");
    }

    try {
      const response = await fetch(`${API_BASE_URL}/admin/baggage`, {
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(
          errorData.error || "Failed to fetch baggage management data"
        );
      }

      const result = await response.json();
      return result.data;
    } catch (error) {
      console.error("Failed to fetch baggage management data:", error);
      throw error;
    }
  },

  async resolveAlert(alertId: string): Promise<void> {
    const token = localStorage.getItem("token");
    if (!token) {
      throw new Error("No authentication token found");
    }

    try {
      const response = await fetch(
        `${API_BASE_URL}/admin/alerts/${alertId}/resolve`,
        {
          method: "PATCH",
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
        }
      );

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Failed to resolve alert");
      }
    } catch (error) {
      console.error("Failed to resolve alert:", error);
      throw error;
    }
  },
};
