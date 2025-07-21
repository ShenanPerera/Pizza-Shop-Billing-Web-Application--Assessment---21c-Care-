import React, { useEffect, useState } from 'react';

interface CartItem {
  pizzaId: number;
  name: string;
  size: string;
  price: number;
  quantity: number;
}

const BillingPage = () => {
  const [cartItems, setCartItems] = useState<CartItem[]>([]);

  useEffect(() => {
    const items: CartItem[] = JSON.parse(localStorage.getItem('cartItems') || '[]');
    setCartItems(items);
  }, []);

  return (
    <div>
      <h1 className="text-3xl font-bold mb-4">Billing Page</h1>
      <ul>
        {cartItems.map((item) => (
          <li key={item.pizzaId}>
            {item.name} - {item.size} - Rs. {item.price} x {item.quantity}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default BillingPage;
