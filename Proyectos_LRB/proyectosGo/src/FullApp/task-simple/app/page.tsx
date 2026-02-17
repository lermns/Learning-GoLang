'use client'

import { useState } from 'react'
import useSWR from 'swr'

const API_URL = 'http://localhost:8080'

interface Tarea {
  id: number
  titulo: string
  descripcion: string
  completado: boolean
}

const fetcher = (url: string) => fetch(url).then((r) => r.json())

export default function Home() {
  const { data: tareas, mutate } = useSWR<Tarea[]>(`${API_URL}/tareas`, fetcher)
  const [titulo, setTitulo] = useState('')
  const [descripcion, setDescripcion] = useState('')

  async function agregarTarea(e: React.FormEvent) {
    e.preventDefault()
    await fetch(`${API_URL}/tareas`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ titulo, descripcion }),
    })
    setTitulo('')
    setDescripcion('')
    mutate()
  }

  async function toggleCompletado(id: number, completado: boolean) {
    await fetch(`${API_URL}/tareas/${id}/completar`, {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ completado: !completado }),
    })
    mutate()
  }

  async function eliminarTarea(id: number) {
    if (confirm('¿Eliminar esta tarea?')) {
      await fetch(`${API_URL}/tareas/${id}`, { method: 'DELETE' })
      mutate()
    }
  }

  return (
    <div className="min-h-screen bg-gray-50 py-8 px-4">
      <div className="max-w-2xl mx-auto">
        <h1 className="text-3xl font-bold mb-8 text-gray-900">📝 Gestor de Tareas</h1>

        {/* Formulario */}
        <form onSubmit={agregarTarea} className="bg-white p-6 rounded-lg shadow mb-8">
          <div className="mb-4">
            <input
              type="text"
              placeholder="Título de la tarea"
              value={titulo}
              onChange={(e) => setTitulo(e.target.value)}
              required
              className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div className="mb-4">
            <textarea
              placeholder="Descripción"
              value={descripcion}
              onChange={(e) => setDescripcion(e.target.value)}
              className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              rows={3}
            />
          </div>
          <button
            type="submit"
            className="w-full bg-blue-500 text-white py-2 rounded-lg hover:bg-blue-600 transition"
          >
            ➕ Agregar Tarea
          </button>
        </form>

        {/* Lista de tareas */}
        <div className="space-y-3">
          {!tareas ? (
            <p className="text-center text-gray-500">Cargando...</p>
          ) : tareas.length === 0 ? (
            <p className="text-center text-gray-500">No hay tareas. ¡Agrega una!</p>
          ) : (
            tareas.map((tarea) => (
              <div
                key={tarea.id}
                className="bg-white p-4 rounded-lg shadow flex items-start gap-3"
              >
                <input
                  type="checkbox"
                  checked={tarea.completado}
                  onChange={() => toggleCompletado(tarea.id, tarea.completado)}
                  className="mt-1 w-5 h-5 cursor-pointer"
                />
                <div className="flex-1">
                  <h3
                    className={`font-semibold ${
                      tarea.completado ? 'line-through text-gray-400' : 'text-gray-900'
                    }`}
                  >
                    {tarea.titulo}
                  </h3>
                  <p className="text-sm text-gray-600 mt-1">{tarea.descripcion}</p>
                </div>
                <button
                  onClick={() => eliminarTarea(tarea.id)}
                  className="text-red-500 hover:text-red-700 px-3 py-1 rounded hover:bg-red-50"
                >
                  🗑️
                </button>
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  )
}
