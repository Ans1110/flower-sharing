interface FlowersDetailProps {
  params: Promise<{
    id: string;
  }>;
}

export default async function FlowersDetail({ params }: FlowersDetailProps) {
  const { id } = await params;

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6">Flower Detail</h1>
      <p>Viewing flower with ID: {id}</p>
    </div>
  );
}
