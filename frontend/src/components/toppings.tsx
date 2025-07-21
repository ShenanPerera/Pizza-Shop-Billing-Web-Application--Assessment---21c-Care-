import React, { useEffect, useState } from 'react';

interface Topping {
  id: number;
  name: string;
  price: number;
  is_active: boolean;
}

interface ToppingsTableProps {
  onNavigateToBilling: () => void;
}

const ToppingsTable: React.FC<ToppingsTableProps> = ({ onNavigateToBilling }) => {
  const [toppings, setToppings] = useState<Topping[]>([]);
  const [selectedToppings, setSelectedToppings] = useState<{ [id: number]: number }>({});

  useEffect(() => {
    const fetchToppings = async () => {
      try {
        console.log('Fetching toppings...');
        const response = await fetch('http://localhost:8080/api/toppings');
        const jsonData = await response.json();
        console.log('Fetched toppings:', jsonData);

        if (response.ok) {
          setToppings(jsonData.data);
        }
      } catch (error) {
        console.error('Failed to fetch toppings:', error);
      }
    };

    fetchToppings();
  }, []);

  const formatPrice = (price: number) => `Rs. ${price.toLocaleString()}`;

  const handleSelectTopping = (toppingId: number, isChecked: boolean) => {
    setSelectedToppings((prev) => {
      if (isChecked) {
        return { ...prev, [toppingId]: prev[toppingId] || 1 };
      } else {
        const updated = { ...prev };
        delete updated[toppingId];
        return updated;
      }
    });
  };

  const handleQuantityChange = (toppingId: number, quantity: number) => {
    setSelectedToppings((prev) => ({
      ...prev,
      [toppingId]: quantity,
    }));
  };

  // const handleAddToCart = () => {
  //   console.log('Adding selected toppings to cart:', selectedToppings);
  //   const cartItems = toppings
  //     .filter((topping) => Object.prototype.hasOwnProperty.call(selectedToppings, topping.id))
  //     .map((topping) => ({
  //       toppingId: topping.id,
  //       name: topping.name,
  //       type: 'topping',
  //       size: 'N/A', 
  //       price: topping.price,
  //       quantity: selectedToppings[topping.id],
  //     }));
  //   console.log('Cart items:', cartItems);

  //   // Save to localStorage - you may want to merge with existing cart if needed
  //   localStorage.setItem('toppingCartItems', JSON.stringify(cartItems));
  //   onNavigateToBilling();
  // };

  const handleAddToCart = () => {
    const existingCart: CartItem[] = JSON.parse(localStorage.getItem('cartItems') || '[]');

    const newItems: CartItem[] = toppings
      .filter((t) => selectedToppings[t.id])
      .map((t) => ({
        id: t.id,
        type: 'topping',
        name: t.name,
        price: t.price,
        quantity: selectedToppings[t.id],
      }));

    // Merge: If topping already exists, update quantity
    const mergedCart = [...existingCart];
    newItems.forEach((newItem) => {
        const existingIndex = mergedCart.findIndex(
            (item) => item.id === newItem.id && item.type === 'topping'
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
      <h1 className="text-3xl font-bold text-gray-800 mb-6">Extra Toppings</h1>

      <div className="bg-white rounded-lg shadow-md overflow-hidden">
        <div className="overflow-x-auto">
          <table className="w-full">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Select</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Price</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Active</th>
                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Quantity</th>
              </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              {toppings.map((topping) => {
                const isSelected = Object.prototype.hasOwnProperty.call(selectedToppings, topping.id);
                return (
                  <tr key={topping.id} className="hover:bg-gray-50">
                    <td className="px-6 py-4 whitespace-nowrap">
                      <input
                        type="checkbox"
                        className="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                        checked={isSelected}
                        onChange={(e) => handleSelectTopping(topping.id, e.target.checked)}
                      />
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      <div className="text-sm font-medium text-gray-900">{topping.name}</div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      <div className="text-sm font-medium text-green-600">{formatPrice(topping.price)}</div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      <span
                        className={`inline-flex px-2 py-1 text-xs font-semibold rounded-full ${
                          topping.is_active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                        }`}
                      >
                        {topping.is_active ? 'Yes' : 'No'}
                      </span>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      <select
                        className="block w-20 px-3 py-1 text-sm border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100 disabled:text-gray-400"
                        value={selectedToppings[topping.id] || 1}
                        onChange={(e) => handleQuantityChange(topping.id, Number(e.target.value))}
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
          disabled={Object.keys(selectedToppings).length === 0}
          className="bg-blue-600 text-white px-4 py-2 rounded-md disabled:opacity-50"
        >
          Add to Cart
        </button>
      </div>
    </div>
  );
};

export default ToppingsTable;
