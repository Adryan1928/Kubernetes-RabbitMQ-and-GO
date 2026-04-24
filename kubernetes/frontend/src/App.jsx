import { useEffect, useMemo, useState } from 'react'

const emptyForm = {
  learner_name: '',
  topic: '',
  level: 'iniciante',
  duration_days: 14,
  goals: '',
}

const statusLabel = {
  pending: 'Na fila',
  processing: 'Processando',
  completed: 'Concluido',
}

function App() {
  const [jobs, setJobs] = useState([])
  const [stats, setStats] = useState(null)
  const [form, setForm] = useState(emptyForm)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  async function loadDashboard() {
    const [jobsResponse, statsResponse] = await Promise.all([
      fetch('/api/jobs'),
      fetch('/api/stats'),
    ])

    if (!jobsResponse.ok || !statsResponse.ok) {
      throw new Error('Nao foi possivel carregar o painel agora.')
    }

    const [jobsData, statsData] = await Promise.all([
      jobsResponse.json(),
      statsResponse.json(),
    ])

    setJobs(jobsData)
    setStats(statsData)
  }

  useEffect(() => {
    loadDashboard().catch((err) => setError(err.message))

    const interval = window.setInterval(() => {
      loadDashboard().catch(() => {})
    }, 5000)

    return () => window.clearInterval(interval)
  }, [])

  async function handleSubmit(event) {
    event.preventDefault()
    setLoading(true)
    setError('')

    try {
      const response = await fetch('/api/jobs', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          ...form,
          duration_days: Number(form.duration_days),
        }),
      })

      if (!response.ok) {
        throw new Error('Verifique os campos e tente novamente.')
      }

      setForm(emptyForm)
      await loadDashboard()
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const completedPlans = useMemo(
    () => jobs.filter((job) => job.status === 'completed').slice(0, 3),
    [jobs],
  )

  return (
    <div className="page-shell">
      <header className="hero">
        <div>
          <p className="eyebrow">Projeto pratico de Kubernetes</p>
          <h1>StudyFlow Lab</h1>
          <p className="hero-copy">
            Um laboratorio de microservicos onde o frontend envia pedidos de planos de estudo,
            a API registra no PostgreSQL, o Redis segura a fila e o worker processa tudo de forma assincrona.
          </p>
        </div>
        <div className="hero-card">
          <span>Fluxo real</span>
          <strong>React - FastAPI - Worker - Redis - PostgreSQL</strong>
          <p>
            Ideal para estudar filas, observabilidade basica, escalonamento e o motivo de separar
            servicos por responsabilidade.
          </p>
        </div>
      </header>

      <main className="content-grid">
        <section className="panel form-panel">
          <div className="panel-heading">
            <h2>Novo pedido</h2>
            <p>Crie uma solicitacao e acompanhe o worker concluir o plano.</p>
          </div>

          <form onSubmit={handleSubmit} className="study-form">
            <label>
              Nome
              <input
                value={form.learner_name}
                onChange={(event) => setForm({ ...form, learner_name: event.target.value })}
                placeholder="Ex.: Adriano"
                required
              />
            </label>

            <label>
              Tema de estudo
              <input
                value={form.topic}
                onChange={(event) => setForm({ ...form, topic: event.target.value })}
                placeholder="Ex.: Kubernetes para iniciantes"
                required
              />
            </label>

            <div className="form-row">
              <label>
                Nivel
                <select
                  value={form.level}
                  onChange={(event) => setForm({ ...form, level: event.target.value })}
                >
                  <option value="iniciante">Iniciante</option>
                  <option value="intermediario">Intermediario</option>
                  <option value="avancado">Avancado</option>
                </select>
              </label>

              <label>
                Dias disponiveis
                <input
                  type="number"
                  min="1"
                  max="90"
                  value={form.duration_days}
                  onChange={(event) => setForm({ ...form, duration_days: event.target.value })}
                  required
                />
              </label>
            </div>

            <label>
              Objetivo
              <textarea
                value={form.goals}
                onChange={(event) => setForm({ ...form, goals: event.target.value })}
                placeholder="Ex.: Quero aprender deploy, services, configmaps e autoscaling na pratica."
                rows="5"
                required
              />
            </label>

            <button disabled={loading} type="submit">
              {loading ? 'Enviando...' : 'Enviar para a fila'}
            </button>
          </form>

          {error ? <p className="error-box">{error}</p> : null}
        </section>

        <section className="side-column">
          <div className="stats-grid">
            <article className="panel stat-card">
              <span>Total</span>
              <strong>{stats?.total_requests ?? '--'}</strong>
            </article>
            <article className="panel stat-card">
              <span>Na fila</span>
              <strong>{stats?.pending_requests ?? '--'}</strong>
            </article>
            <article className="panel stat-card">
              <span>Em processo</span>
              <strong>{stats?.processing_requests ?? '--'}</strong>
            </article>
            <article className="panel stat-card accent-card">
              <span>Concluidos</span>
              <strong>{stats?.completed_requests ?? '--'}</strong>
            </article>
          </div>

          <div className="panel insights-panel">
            <div className="panel-heading">
              <h2>Planos prontos</h2>
              <p>Esses cards mostram o resultado persistido no PostgreSQL.</p>
            </div>
            <div className="insights-list">
              {completedPlans.length === 0 ? (
                <p className="empty-state">Assim que o worker finalizar pedidos, eles aparecem aqui.</p>
              ) : (
                completedPlans.map((job) => (
                  <article key={job.id} className="insight-card">
                    <header>
                      <strong>{job.topic}</strong>
                      <span>{job.learner_name}</span>
                    </header>
                    <p>{job.generated_plan}</p>
                  </article>
                ))
              )}
            </div>
          </div>
        </section>
      </main>

      <section className="panel jobs-panel">
        <div className="panel-heading jobs-heading">
          <div>
            <h2>Fila de processamento</h2>
            <p>Cada linha representa um job criado pela API e acompanhado pelo frontend.</p>
          </div>
          <span className="refresh-chip">Atualizacao automatica a cada 5s</span>
        </div>

        <div className="jobs-list">
          {jobs.length === 0 ? (
            <p className="empty-state">Nenhum pedido ainda. Crie o primeiro para testar o fluxo completo.</p>
          ) : (
            jobs.map((job) => (
              <article key={job.id} className="job-card">
                <div>
                  <p className="job-title">{job.topic}</p>
                  <p className="job-meta">
                    {job.learner_name} - {job.level} - {job.duration_days} dias
                  </p>
                </div>
                <span className={`status-pill status-${job.status}`}>{statusLabel[job.status]}</span>
              </article>
            ))
          )}
        </div>
      </section>
    </div>
  )
}

export default App
