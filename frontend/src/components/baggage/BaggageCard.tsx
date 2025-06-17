import { Package, AlertTriangle, CheckCircle, Truck } from 'lucide-react';
import type { BaggageItem } from '@/types';

interface BaggageCardProps {
  baggage: BaggageItem;
}

export function BaggageCard({ baggage }: BaggageCardProps) {
  const getStatusIcon = (status: BaggageItem['status']) => {
    switch (status) {
      case 'CHECKED':
        return <CheckCircle className="h-5 w-5 text-blue-500" />;
      case 'IN_TRANSIT':
        return <Truck className="h-5 w-5 text-yellow-500" />;
      case 'DELIVERED':
        return <CheckCircle className="h-5 w-5 text-green-600" />;
      case 'LOST':
        return <AlertTriangle className="h-5 w-5 text-red-500" />;
      default:
        return <Package className="h-5 w-5 text-gray-500" />;
    }
  };

  const getStatusColor = (status: BaggageItem['status']) => {
    switch (status) {
      case 'CHECKED':
        return 'bg-blue-100 text-blue-800 border-blue-200';
      case 'IN_TRANSIT':
        return 'bg-yellow-100 text-yellow-800 border-yellow-200';
      case 'DELIVERED':
        return 'bg-green-100 text-green-800 border-green-200';
      case 'LOST':
        return 'bg-red-100 text-red-800 border-red-200';
      default:
        return 'bg-gray-100 text-gray-800 border-gray-200';
    }
  };

  const formatWeight = (weight?: number) => {
    if (!weight) return 'N/A';
    return `${weight}kg`;
  };

  return (
    <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow">
      <div className="flex items-start justify-between mb-4">
        <div className="flex items-center space-x-3">
          <div className="p-2 bg-blue-50 rounded-lg">
            <Package className="h-6 w-6 text-blue-600" />
          </div>
          <div>
            {baggage.trackingNumber && (
              <h3 className="font-semibold text-gray-900">{baggage.trackingNumber}</h3>
            )}
          </div>
        </div>
        <span className={`inline-flex items-center gap-1 px-3 py-1 rounded-full text-sm font-medium border ${getStatusColor(baggage.status)}`}>
          {getStatusIcon(baggage.status)}
          {baggage.status.replaceAll('_', ' ')}
        </span>
      </div>

      <div className="grid grid-cols-2 gap-4 mb-4">
        <div>
          <p className="text-sm font-medium text-gray-500">Weight</p>
          <p className="text-gray-900">{formatWeight(baggage.weight)}</p>
        </div>
      </div>

      {baggage.specialHandling && baggage.specialHandling.length > 0 && (
        <div className="border-t pt-4">
          <p className="text-sm font-medium text-gray-500 mb-2">Special Handling</p>
          <div className="flex flex-wrap gap-2">
            {baggage.specialHandling.split(" ").map((handling, index) => (
              <span
                key={index}
                className="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium bg-purple-100 text-purple-800"
              >
                {handling}
              </span>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
