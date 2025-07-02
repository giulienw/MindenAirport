import { useEffect, useState } from "react";
import { Link } from "react-router";
import { FlightBoard } from "@/components/flight";
import { useFlights } from "@/hooks";
import { Plane, Clock, MapPin, Users } from "lucide-react";

function Home() {
  const { flights, loading, error, refetch } = useFlights();
  const [currentTime, setCurrentTime] = useState(new Date());
  const isAdmin = localStorage.getItem('user') ? JSON.parse(localStorage.getItem('user')!).role === 'ADMIN' : false;

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentTime(new Date());
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  const stats = {
    totalFlights: flights.length,
    onTime: flights.filter(f => f.statusInfo?.name?.toLowerCase() === 'on time').length,
    delayed: flights.filter(f => f.statusInfo?.name?.toLowerCase() === 'delayed').length,
    departures: flights.filter(f => {
      const now = new Date();
      const departure = new Date(f.scheduledDeparture);
      return departure >= now;
    }).length,
  };

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Airport Header */}
      <div className="bg-gradient-to-r from-blue-600 to-blue-800 text-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <div className="flex flex-col md:flex-row md:items-center md:justify-between">
            <div>
              <h1 className="text-4xl font-bold mb-2">Minden Airport</h1>
              <p className="text-blue-100 text-lg">
                Welcome to Minden Airport - Your Gateway to the World
              </p>
            </div>
            <div className="mt-4 md:mt-0 flex flex-col items-end space-y-2">
              <div className="text-right">
                <div className="text-2xl font-mono">
                  {currentTime.toLocaleTimeString('en-US', {
                    hour: '2-digit',
                    minute: '2-digit',
                    second: '2-digit',
                    hour12: false,
                  })}
                </div>
                <div className="text-blue-100">
                  {currentTime.toLocaleDateString('en-US', {
                    weekday: 'long',
                    year: 'numeric',
                    month: 'long',
                    day: 'numeric',
                  })}
                </div>
              </div>
              {isAdmin ? (
                <Link 
                to="/admin" 
                className="bg-white text-red-600 hover:bg-red-50 px-4 py-2 rounded-md text-sm font-medium transition-colors"
              >
                Admin
              </Link>
              ) : null}
              <Link 
                to="/login" 
                className="bg-white text-blue-600 hover:bg-blue-50 px-4 py-2 rounded-md text-sm font-medium transition-colors"
              >
                My Tickets
              </Link>
            </div>
          </div>
        </div>
      </div>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Stats Section */}
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4 mb-8">
          <div className="bg-white rounded-lg shadow p-6 text-center">
            <div className="flex justify-center mb-2">
              <Plane className="h-8 w-8 text-blue-600" />
            </div>
            <div className="text-2xl font-bold text-gray-900">{stats.totalFlights}</div>
            <div className="text-sm text-gray-600">Total Flights</div>
          </div>
          <div className="bg-white rounded-lg shadow p-6 text-center">
            <div className="flex justify-center mb-2">
              <Clock className="h-8 w-8 text-green-600" />
            </div>
            <div className="text-2xl font-bold text-gray-900">{stats.onTime}</div>
            <div className="text-sm text-gray-600">On Time</div>
          </div>
          <div className="bg-white rounded-lg shadow p-6 text-center">
            <div className="flex justify-center mb-2">
              <MapPin className="h-8 w-8 text-yellow-600" />
            </div>
            <div className="text-2xl font-bold text-gray-900">{stats.delayed}</div>
            <div className="text-sm text-gray-600">Delayed</div>
          </div>
          <div className="bg-white rounded-lg shadow p-6 text-center">
            <div className="flex justify-center mb-2">
              <Users className="h-8 w-8 text-purple-600" />
            </div>
            <div className="text-2xl font-bold text-gray-900">{stats.departures}</div>
            <div className="text-sm text-gray-600">Departures Today</div>
          </div>
        </div>

        {/* Flight Information Section */}
        <div className="mb-8">
          <div className="flex items-center justify-between mb-6">
            <h2 className="text-3xl font-bold text-gray-900">Flight Information</h2>
            <button
              onClick={refetch}
              className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg transition-colors duration-200 flex items-center space-x-2"
              disabled={loading}
            >
              <Clock className="h-4 w-4" />
              <span>{loading ? 'Refreshing...' : 'Refresh'}</span>
            </button>
          </div>

          {error ? (
            <div className="bg-red-50 border border-red-200 rounded-lg p-6">
              <div className="flex items-center">
                <div className="text-red-600">
                  <svg className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div className="ml-3">
                  <h3 className="text-sm font-medium text-red-800">Error Loading Flight Data</h3>
                  <p className="text-sm text-red-700 mt-1">{error}</p>
                  <button
                    onClick={refetch}
                    className="text-sm text-red-800 underline hover:text-red-900 mt-2"
                  >
                    Try again
                  </button>
                </div>
              </div>
            </div>
          ) : (
            <FlightBoard flights={flights} loading={loading} />
          )}
        </div>
      </div>
    </div>
  );
}

export default Home;
