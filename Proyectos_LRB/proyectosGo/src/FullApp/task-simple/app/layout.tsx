import './globals.css'
import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Gestor de Tareas',
  description: 'App de gestión de tareas',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="es">
      <body>{children}</body>
    </html>
  )
}
