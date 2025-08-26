import {useEffect, useState } from "react";

function BooksTable() {
    const [books, setBooks] = useState([]);

    useEffect(() => {
        fetch("http://localhost:8000/books")
            .then((res) => res.json())
            .then((data) => setBooks(data))
            .catch((err) => console.error("Error fetching books: ", err));
    }, []);

    return (
        <div className="p-4">
            <h1 className="text-x1 font-bold mb-4">Books</h1>
            <table className="table-auto border-collapse border border-gray-400 w-full">
                <thead>
                    <tr>
                        <th className="border border-gray-400 px-2 py-1">ID</th>
                        <th className="border border-gray-400 px-2 py-1">Title</th>
                        <th className="border border-gray-400 px-2 py-1">Author</th>
                        <th className="border border-gray-400 px-2 py-1">Date Bought</th>
                        <th className="border border-gray-400 px-2 py-1">Status</th>
                        <th className="border border-gray-400 px-2 py-1">Date Read</th>
                        <th className="border border-gray-400 px-2 py-1">Description</th>
                    </tr>
                    
                </thead>
                <tbody>
                    {books.map((book) => (
                        <tr key={book.id}>
                        <td className="border border-gray-400 px-2 py-1">{book.id}</td>
                        <td className="border border-gray-400 px-2 py-1">{book.title}</td>
                        <td className="border border-gray-400 px-2 py-1">{book.author}</td>
                        <td className="border border-gray-400 px-2 py-1">
                            {new Date(book.date_bought).toLocaleDateString()}
                        </td>
                        <td className="border border-gray-400 px-2 py-1">
                            {book.status === 1 ? "Read" : "Unread"}
                        </td>
                        <td className="border border-gray-400 px-2 py-1">
                            {book.date_read ? new Date(book.date_read).toLocaleDateString() : "-"}
                        </td>
                        <td className="border border-gray-400 px-2 py-1">{book.description}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}

export default BooksTable;