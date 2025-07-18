/**
 * BaggageView Component
 * 
 * A comprehensive baggage management interface that displays user's baggage
 * with statistics, filtering, and management capabilities. Provides a dashboard
 * view of all baggage items with status-based organization.
 * 
 * Features:
 * - Baggage statistics dashboard (total, checked, carry-on, in-transit)
 * - Search functionality for finding specific items
 * - Status and type-based filtering capabilities
 * - Responsive grid layout for baggage cards
 * - Manual refresh functionality
 * - Export capabilities for baggage data
 * 
 * @param props - Component props
 * @param props.baggage - Array of baggage items to display
 * @param props.onRefresh - Callback function for refresh operations
 * @param props.loading - Loading state for UI feedback
 */

import { useState } from 'react';
import { Package, Search, Download, RefreshCw } from 'lucide-react';
import type { BaggageItem } from '@/types';
import { BaggageCard } from './BaggageCard';
import { Button } from '@/components/ui';

interface BaggageViewProps {
  /** Array of baggage items to display */
  baggage: BaggageItem[];
  /** Callback function for manual refresh operations */
  onRefresh?: () => void;
  /** Loading state for showing UI feedback */
  loading?: boolean;
}

export function BaggageView({ baggage, onRefresh, loading = false }: BaggageViewProps) {
  // Search term state for filtering baggage items by tracking number
  const [searchTerm, setSearchTerm] = useState('');
  // Status filter for showing specific baggage statuses
  const [statusFilter, setStatusFilter] = useState<string>('all');
  // Type filter for showing specific baggage types (carry-on, checked, etc.)
  const [typeFilter, setTypeFilter] = useState<string>('all');

  /**
   * Filter baggage items based on search term, status, and type filters
   */
  const filteredBaggage = baggage.filter(item => {
    const matchesSearch = !searchTerm || 
      item.trackingNumber?.toLowerCase().includes(searchTerm.toLowerCase());
    
    const matchesStatus = statusFilter === 'all' || item.status === statusFilter;
    const matchesType = typeFilter === 'all' || item.type === typeFilter;
    
    return matchesSearch && matchesStatus && matchesType;
  });

  /**
   * Calculate statistics for baggage dashboard
   * Groups baggage by status for display metrics
   */
  const getStatusStats = () => {
    const stats = baggage.reduce((acc, item) => {
      acc[item.status] = (acc[item.status] || 0) + 1;
      return acc;
    }, {} as Record<string, number>);
    return stats;
  };

  const statusStats = getStatusStats();

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div className="flex items-center space-x-3">
          <div className="p-2 bg-blue-50 rounded-lg">
            <Package className="h-6 w-6 text-blue-600" />
          </div>
          <div>
            <h2 className="text-2xl font-bold text-gray-900">My Baggage</h2>
            <p className="text-gray-600">Track and manage your luggage</p>
          </div>
        </div>
        <div className="flex space-x-3">
          <Button variant="outline" onClick={onRefresh} disabled={loading}>
            <RefreshCw className={`h-4 w-4 mr-2 ${loading ? 'animate-spin' : ''}`} />
            Refresh
          </Button>
          <Button variant="outline">
            <Download className="h-4 w-4 mr-2" />
            Export
          </Button>
        </div>
      </div>

      {/* Stats */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm font-medium text-gray-500">Total Items</p>
              <p className="text-2xl font-bold text-gray-900">{baggage.length}</p>
            </div>
            <Package className="h-8 w-8 text-blue-500" />
          </div>
        </div>
        <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm font-medium text-gray-500">Checked Bags</p>
              <p className="text-2xl font-bold text-gray-900">{baggage.filter(item => item.status == "CHECKED").length}</p>
            </div>
            <Package className="h-8 w-8 text-green-500" />
          </div>
        </div>
        <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm font-medium text-gray-500">Carry-on</p>
              <p className="text-2xl font-bold text-gray-900">{baggage.filter(item => item.type == "CARRY_ON").length}</p>
            </div>
            <Package className="h-8 w-8 text-orange-500" />
          </div>
        </div>
        <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm font-medium text-gray-500">In Transit</p>
              <p className="text-2xl font-bold text-gray-900">
                {statusStats['IN_TRANSIT'] || 0}
              </p>
            </div>
            <Package className="h-8 w-8 text-yellow-500" />
          </div>
        </div>
      </div>

      {/* Filters */}
      <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <div className="flex flex-col lg:flex-row gap-4">
          <div className="flex-1">
            <label htmlFor="search" className="block text-sm font-medium text-gray-700 mb-2">
              Search baggage
            </label>
            <div className="relative">
              <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-gray-400" />
              <input
                id="search"
                type="text"
                placeholder="Search by tracking number, description, or location..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className="pl-10 w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>
          <div className="lg:w-48">
            <label htmlFor="status-filter" className="block text-sm font-medium text-gray-700 mb-2">
              Status
            </label>
            <select
              id="status-filter"
              value={statusFilter}
              onChange={(e) => setStatusFilter(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="all">All Status</option>
              <option value="CHECKED">Checked In</option>
              <option value="IN_TRANSIT">In Transit</option>
              <option value="DELIVERED">Delivered</option>
              <option value="LOST">Lost</option>
            </select>
          </div>
          <div className="lg:w-48">
            <label htmlFor="type-filter" className="block text-sm font-medium text-gray-700 mb-2">
              Type
            </label>
            <select
              id="type-filter"
              value={typeFilter}
              onChange={(e) => setTypeFilter(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
            >
              <option value="all">All Types</option>
              <option value="CHECKED">Checked Bag</option>
              <option value="CARRY_ON">Carry-on</option>
              <option value="PERSONAL">Personal Item</option>
            </select>
          </div>
        </div>
      </div>

      {/* Baggage List */}
      {filteredBaggage.length === 0 ? (
        <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-12 text-center">
          <Package className="h-12 w-12 text-gray-400 mx-auto mb-4" />
          <h3 className="text-lg font-medium text-gray-900 mb-2">No baggage found</h3>
          <p className="text-gray-500">
            {baggage.length === 0
              ? "You don't have any baggage items yet."
              : "No baggage matches your current filters."}
          </p>
        </div>
      ) : (
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          {filteredBaggage.map((item) => (
            <BaggageCard key={item.id} baggage={item} />
          ))}
        </div>
      )}
    </div>
  );
}
