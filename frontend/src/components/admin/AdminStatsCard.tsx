import React from 'react';
import { type LucideIcon, TrendingUp, TrendingDown } from 'lucide-react';

interface AdminStatsCardProps {
  title: string;
  value: string | number;
  icon: LucideIcon;
  color: 'blue' | 'green' | 'purple' | 'emerald' | 'orange' | 'red';
  change?: string;
  subtitle?: string;
}

export const AdminStatsCard: React.FC<AdminStatsCardProps> = ({
  title,
  value,
  icon: Icon,
  color,
  change,
  subtitle,
}) => {
  const colorClasses = {
    blue: 'bg-blue-50 text-blue-600',
    green: 'bg-green-50 text-green-600',
    purple: 'bg-purple-50 text-purple-600',
    emerald: 'bg-emerald-50 text-emerald-600',
    orange: 'bg-orange-50 text-orange-600',
    red: 'bg-red-50 text-red-600',
  };

  const isPositive = change?.startsWith('+');
  const isNegative = change?.startsWith('-');

  return (
    <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow">
      <div className="flex items-center justify-between">
        <div className="flex-1">
          <p className="text-sm font-medium text-gray-600 mb-1">{title}</p>
          <p className="text-3xl font-bold text-gray-900">{value}</p>
          {subtitle && (
            <p className="text-sm text-gray-500 mt-1">{subtitle}</p>
          )}
        </div>
        <div className={`p-3 rounded-lg ${colorClasses[color]}`}>
          <Icon className="w-6 h-6" />
        </div>
      </div>
      
      {change && (
        <div className="mt-4 flex items-center">
          {isPositive && <TrendingUp className="w-4 h-4 text-green-500 mr-1" />}
          {isNegative && <TrendingDown className="w-4 h-4 text-red-500 mr-1" />}
          <span
            className={`text-sm font-medium ${
              isPositive ? 'text-green-600' : isNegative ? 'text-red-600' : 'text-gray-600'
            }`}
          >
            {change}
          </span>
          <span className="text-sm text-gray-500 ml-1">from last month</span>
        </div>
      )}
    </div>
  );
};
