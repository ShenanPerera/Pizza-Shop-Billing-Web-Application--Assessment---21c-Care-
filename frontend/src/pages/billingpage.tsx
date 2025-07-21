import React, { useRef, useState, useEffect } from 'react';
import printJS from 'print-js';

interface CartItem {
  id: number | undefined;
  type: 'pizza' | 'topping' | 'beverage';
  name: string;
  size?: string;
  price: number;
  quantity: number;
}

const BillingPage: React.FC = () => {
  const [cartItems, setCartItems] = useState<CartItem[]>([]);
  const [customerName, setCustomerName] = useState<string>('');
  const [telephone, setTelephone] = useState<string>('');
  const [amountPaid, setAmountPaid] = useState<number>(0);
  const [isPaymentDone, setIsPaymentDone] = useState<boolean>(false);
  const [printDateTime, setPrintDateTime] = useState<Date>(new Date());

  const invoiceRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const items: CartItem[] = JSON.parse(localStorage.getItem('cartItems') || '[]');
    setCartItems(items);
  }, []);

  const formatPrice = (price: number): string => `Rs. ${price.toLocaleString()}`;
  const subtotal: number = cartItems.reduce((sum, item) => sum + item.price * item.quantity, 0);
  const tax: number = subtotal * 0.1; // 10% tax
  const total: number = subtotal + tax;
  const balance: number = amountPaid - total;

  const markPaymentDone = (): void => {
    if (amountPaid >= total) {
      setIsPaymentDone(true);
      console.log('Payment Successful');
    } else {
      alert(`Amount paid (Rs. ${amountPaid.toLocaleString()}) is less than the total (Rs. ${total.toLocaleString()}).`);
    }
  };

  const handlePrintWithPrintJS = () => {
    if (invoiceRef.current) {
      setPrintDateTime(new Date()); // update date/time right before printing
      setTimeout(() => {
        printJS({
          printable: 'invoice-print-area',
          type: 'html',
          header: '<h2>Pizza Palace Invoice</h2>',
          targetStyles: ['*'],
          scanStyles: false,
        });
        console.log('Print initiated with print-js!');
      }, 100); // slight delay so state updates
    } else {
      alert('Invoice content not ready for printing. Please add items.');
    }
  };

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-4">Billing Page</h1>

      <input
        placeholder="Customer Name"
        value={customerName}
        onChange={(e) => setCustomerName(e.target.value)}
        className="border p-2 mb-2 w-full"
      />
      <input
        placeholder="Telephone Number"
        value={telephone}
        onChange={(e) => setTelephone(e.target.value)}
        className="border p-2 mb-2 w-full"
      />

      {cartItems.length === 0 ? (
        <p className="text-lg text-gray-600">Your cart is empty. Please add items to generate an invoice.</p>
      ) : (
        <>
          <div
            id="invoice-print-area"
            ref={invoiceRef}
            style={{ padding: 20, border: '1px solid #ccc', marginBottom: 10, backgroundColor: 'white' }}
            className="invoice-printable-area"
          >
            <h2 className="text-2xl font-semibold mb-3">Invoice</h2>
            <p className="mb-1"><strong>Customer:</strong> {customerName || 'N/A'}</p>
            <p className="mb-3"><strong>Telephone:</strong> {telephone || 'N/A'}</p>
            <p className="mb-3"><strong>Date:</strong> {printDateTime.toLocaleDateString('en-LK', { year: 'numeric', month: 'long', day: 'numeric' })}</p>
            <p className="mb-3"><strong>Time:</strong> {printDateTime.toLocaleTimeString('en-LK')}</p>

            <div className="border-t border-b py-2 my-2">
              <div className="flex justify-between font-bold pb-1">
                <span>Item</span>
                <span>Quantity x Price = Total</span>
              </div>
              {cartItems.map((item, idx) => (
                <div key={`${item.type}-${item.id ?? idx}`} className="flex justify-between text-sm py-0.5">
                  <span>
                    {item.name} {item.size ? `(${item.size})` : ''}
                  </span>
                  <span>
                    {formatPrice(item.price)} x {item.quantity} = {formatPrice(item.price * item.quantity)}
                  </span>
                </div>
              ))}
            </div>

            <hr className="my-3" />

            <div className="text-right">
              <p className="mb-1">Subtotal: <span className="font-medium">{formatPrice(subtotal)}</span></p>
              <p className="mb-1">Tax (10%): <span className="font-medium">{formatPrice(tax)}</span></p>
              <p className="text-xl font-bold mt-2">
                Total: {formatPrice(total)}
              </p>
              <p className="mt-3">Amount Paid: <span className="font-bold">{formatPrice(amountPaid)}</span></p>
              <p className="mt-1">Balance: <span className="font-bold">{formatPrice(balance)}</span></p>
              {isPaymentDone && balance >= 0 && (
                <p className="mt-4 text-center text-lg font-semibold">Thank you for your business!</p>
              )}
            </div>
          </div>

          <input
            type="number"
            placeholder="Amount Paid"
            value={amountPaid === 0 ? '' : amountPaid}
            onChange={(e) => setAmountPaid(Number(e.target.value))}
            className="border p-2 mb-2 w-full"
          />

          <button
            onClick={markPaymentDone}
            disabled={isPaymentDone || amountPaid < total}
            className="bg-green-600 text-white px-4 py-2 rounded mb-2 mr-2 disabled:opacity-50 hover:bg-green-700 transition duration-300"
          >
            {isPaymentDone ? 'Payment Completed' : 'Mark Payment Done'}
          </button>

          <button
            onClick={handlePrintWithPrintJS}
            disabled={!isPaymentDone || cartItems.length === 0 || !invoiceRef.current}
            className="bg-blue-600 text-white px-4 py-2 rounded disabled:opacity-50 hover:bg-blue-700 transition duration-300"
          >
            Print Invoice
          </button>
        </>
      )}
    </div>
  );
};

export default BillingPage;
