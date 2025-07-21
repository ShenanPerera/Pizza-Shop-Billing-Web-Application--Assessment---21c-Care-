import React, { useEffect, useState } from 'react';

interface Pizza {
  id: number;
  item_id: number;
  name: string;
  size: string;
  base_type: string;
  price: number;
  is_active: boolean;
  created_at: string;
  updated_at: string;
  deleted_at: string | null;
}

interface PizzaPageProps {
  onNavigateToBilling: () => void;
}

const Pizza: React.FC<PizzaPageProps> = ({ onNavigateToBilling }) => {
  const [pizzas, setPizzas] = useState<Pizza[]>([]);
  const [selectedPizzas, setSelectedPizzas] = useState<{ [id: number]: number }>({});

  useEffect(() => {
    const fetchPizzas = async () => {
      const response = await fetch('http://localhost:8080/api/pizzas');
      const jsonData = await response.json();
      if (response.ok) {
        setPizzas(jsonData.data);
      }
    };
    fetchPizzas();
  }, []);

  const formatPrice = (price: number) => `Rs. ${price.toLocaleString()}`;

  const handleSelectPizza = (pizzaId: number, isChecked: boolean) => {
    setSelectedPizzas((prev) => {
      if (isChecked) {
        return { ...prev, [pizzaId]: prev[pizzaId] || 1 };
      } else {
        const updated = { ...prev };
        delete updated[pizzaId];
        return updated;
      }
    });
  };

  const handleQuantityChange = (pizzaId: number, quantity: number) => {
    setSelectedPizzas((prev) => ({
      ...prev,
      [pizzaId]: quantity,
    }));
  };

  const handleAddToCart = () => {
    const cartItems = pizzas
      .filter((pizza) => Object.prototype.hasOwnProperty.call(selectedPizzas, pizza.id))
      .map((pizza) => ({
        pizzaId: pizza.id,
        name: pizza.name,
        type: 'pizza',
        size: pizza.size,
        price: pizza.price,
        quantity: selectedPizzas[pizza.id],
      }));

    localStorage.setItem('cartItems', JSON.stringify(cartItems));
    onNavigateToBilling();
  };

  return (
    <>
      <div className="p-6">
        <h1 className="text-3xl font-bold text-gray-800 mb-6">Pizza Menu</h1>

        <div className="bg-white rounded-lg shadow-md overflow-hidden">
          <div className="overflow-x-auto">
            <table className="w-full">
              <thead className="bg-gray-50">
                <tr>
                  <th className="px-6 py-3">Select</th>
                  <th className="px-6 py-3">Name</th>
                  <th className="px-6 py-3">Size</th>
                  <th className="px-6 py-3">Base Type</th>
                  <th className="px-6 py-3">Price</th>
                  <th className="px-6 py-3">Active</th>
                  <th className="px-6 py-3">Quantity</th>
                </tr>
              </thead>
              <tbody>
                {pizzas.map((pizza) => {
                  const isSelected = Object.prototype.hasOwnProperty.call(selectedPizzas, pizza.id);
                  return (
                    <tr key={pizza.id} className="hover:bg-gray-50">
                      <td className="px-6 py-4">
                        <input
                          type="checkbox"
                          checked={isSelected}
                          onChange={(e) => handleSelectPizza(pizza.id, e.target.checked)}
                          className="h-4 w-4 text-blue-600"
                        />
                      </td>
                      <td className="px-6 py-4">{pizza.name}</td>
                      <td className="px-6 py-4">{pizza.size}</td>
                      <td className="px-6 py-4">{pizza.base_type}</td>
                      <td className="px-6 py-4">{formatPrice(pizza.price)}</td>
                      <td className="px-6 py-4">
                        <span
                          className={`px-2 py-1 rounded-full text-xs ${
                            pizza.is_active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                          }`}
                        >
                          {pizza.is_active ? 'Yes' : 'No'}
                        </span>
                      </td>
                      <td className="px-6 py-4">
                        <select
                          className="w-20 px-3 py-1 text-sm border rounded-md"
                          value={selectedPizzas[pizza.id] || 1}
                          onChange={(e) => handleQuantityChange(pizza.id, Number(e.target.value))}
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
      </div>
      <div>
        <button
          onClick={handleAddToCart}
          className="bg-blue-600 text-white px-4 py-2 rounded-md"
          disabled={Object.keys(selectedPizzas).length === 0}
        >
          Add to Cart
        </button>
      </div>
    </>
  );
};

export default Pizza;
