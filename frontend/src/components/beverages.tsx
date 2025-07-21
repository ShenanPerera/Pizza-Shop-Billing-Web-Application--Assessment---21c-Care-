import React, { useEffect, useState } from 'react';

interface Beverage {
  id: number;
  name: string;
  price: number;
  size: string;
  is_active: boolean;
}

interface BeveragesTableProps {
  onNavigateToBilling: () => void;
}

interface CartItem {
  id: number;
  type: 'pizza' | 'topping' | 'beverage';
  name: string;
  size?: string;
  price: number;
  quantity: number;
}

const BeveragesTable: React.FC<BeveragesTableProps> = ({ onNavigateToBilling }) => {
  const [beverages, setBeverages] = useState<Beverage[]>([]);
  const [selectedBeverages, setSelectedBeverages] = useState<{ [id: number]: number }>({});

  useEffect(() => {
    const fetchBeverages = async () => {
      try {
        console.log('Fetching beverages...');
        const response = await fetch('http://localhost:8080/api/beverages');
        const jsonData = await response.json();
        console.log('Fetched beverages:', jsonData);

        if (response.ok) {
          setBeverages(jsonData.data);
        }
      } catch (error) {
        console.error('Failed to fetch beverages:', error);
      }
    };

    fetchBeverages();
  }, []);

  const formatPrice = (price: number) => `Rs. ${price.toLocaleString()}`;

  const handleSelectBeverage = (beverageId: number, isChecked: boolean) => {
    setSelectedBeverages((prev) => {
      if (isChecked) {
        return { ...prev, [beverageId]: prev[beverageId] || 1 };
      } else {
        const updated = { ...prev };
        delete updated[beverageId];
        return updated;
      }
    });
  };

  const handleQuantityChange = (beverageId: number, quantity: number) => {
    setSelectedBeverages((prev) => ({
      ...prev,
      [beverageId]: quantity,
    }));
  };

  const handleAddToCart = () => {
    const existingCart: CartItem[] = JSON.parse(localStorage.getItem('cartItems') || '[]');

    const newItems: CartItem[] = beverages
      .filter((b) => selectedBeverages[b.id])
      .map((b) => ({
        id: b.id,
        type: 'beverage',
        name: b.name,
        size: b.size,
        price: b.price,
        quantity: selectedBeverages[b.id],
      }));

    const mergedCart = [...existingCart];
    newItems.forEach((newItem) => {
      const existingIndex = mergedCart.findIndex(
        (item) => item.id === newItem.id && item.type === 'beverage'
      );
      if (existingIndex > -1) {
        mergedCart[existingIndex].quantity += newItem.quantity;
      } else {
        mergedCart.push(newItem);
      }
    });

    localStorage.setItem('cartItems', JSON.stringify(mergedCart));
    onNavigateToBilling();
  };

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold text-gray-800 mb-6">Refreshing Beverages</h1>

      <div className="bg-white rounded-lg shadow-md overflow-hidden">
        <div className="overflow-x-auto">
          <table className="w-full">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3">Select</th>
                <th className="px-6 py-3">Name</th>
                <th className="px-6 py-3">Size</th>
                <th className="px-6 py-3">Price</th>
                <th className="px-6 py-3">Active</th>
                <th className="px-6 py-3">Quantity</th>
              </tr>
            </thead>
            <tbody>
              {beverages.map((beverage) => {
                const isSelected = Object.prototype.hasOwnProperty.call(selectedBeverages, beverage.id);
                return (
                  <tr key={beverage.id} className="hover:bg-gray-50">
                    <td className="px-6 py-4">
                      <input
                        type="checkbox"
                        className="h-4 w-4 text-blue-600"
                        checked={isSelected}
                        onChange={(e) => handleSelectBeverage(beverage.id, e.target.checked)}
                      />
                    </td>
                    <td className="px-6 py-4">{beverage.name}</td>
                    <td className="px-6 py-4">{beverage.size}</td>
                    <td className="px-6 py-4">{formatPrice(beverage.price)}</td>
                    <td className="px-6 py-4">
                      <span
                        className={`inline-flex px-2 py-1 text-xs font-semibold rounded-full ${
                          beverage.is_active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                        }`}
                      >
                        {beverage.is_active ? 'Yes' : 'No'}
                      </span>
                    </td>
                    <td className="px-6 py-4">
                      <select
                        className="w-20 px-3 py-1 text-sm border rounded-md"
                        value={selectedBeverages[beverage.id] || 1}
                        onChange={(e) => handleQuantityChange(beverage.id, Number(e.target.value))}
                        disabled={!isSelected}
                      >
                        {[...Array(10)].map((_, i) => (
                          <option key={i + 1} value={i + 1}>
                            {i + 1}
                          </option>
                        ))}
                      </select>
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </table>
        </div>
      </div>

      <div className="mt-4">
        <button
          onClick={handleAddToCart}
          disabled={Object.keys(selectedBeverages).length === 0}
          className="bg-blue-600 text-white px-4 py-2 rounded-md disabled:opacity-50"
        >
          Add to Cart
        </button>
      </div>
    </div>
  );
};

export default BeveragesTable;
