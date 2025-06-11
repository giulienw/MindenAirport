import { useState, useCallback } from 'react'

interface UseFormOptions<T> {
  initialValues: T
  validate?: (values: T) => Partial<Record<keyof T, string>>
  onSubmit?: (values: T) => void | Promise<void>
}

interface UseFormReturn<T> {
  values: T
  errors: Partial<Record<keyof T, string>>
  isSubmitting: boolean
  handleChange: (e: React.ChangeEvent<HTMLInputElement>) => void
  handleSubmit: (e: React.FormEvent) => void
  setFieldValue: (name: keyof T, value: T[keyof T]) => void
  setFieldError: (name: keyof T, error: string) => void
  resetForm: () => void
}

export function useForm<T extends Record<string, unknown>>({
  initialValues,
  validate,
  onSubmit,
}: UseFormOptions<T>): UseFormReturn<T> {
  const [values, setValues] = useState<T>(initialValues)
  const [errors, setErrors] = useState<Partial<Record<keyof T, string>>>({})
  const [isSubmitting, setIsSubmitting] = useState(false)

  const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = e.target
    const fieldValue = type === 'checkbox' ? checked : value

    setValues(prev => ({
      ...prev,
      [name]: fieldValue,
    }))

    // Clear error when user starts typing
    if (errors[name as keyof T]) {
      setErrors(prev => ({
        ...prev,
        [name]: undefined,
      }))
    }
  }, [errors])

  const setFieldValue = useCallback((name: keyof T, value: T[keyof T]) => {
    setValues(prev => ({
      ...prev,
      [name]: value,
    }))
  }, [])

  const setFieldError = useCallback((name: keyof T, error: string) => {
    setErrors(prev => ({
      ...prev,
      [name]: error,
    }))
  }, [])

  const handleSubmit = useCallback(async (e: React.FormEvent) => {
    e.preventDefault()
    
    if (validate) {
      const validationErrors = validate(values)
      setErrors(validationErrors)
      
      if (Object.keys(validationErrors).length > 0) {
        return
      }
    }

    if (onSubmit) {
      setIsSubmitting(true)
      try {
        await onSubmit(values)
      } catch (error) {
        console.error('Form submission error:', error)
      } finally {
        setIsSubmitting(false)
      }
    }
  }, [values, validate, onSubmit])

  const resetForm = useCallback(() => {
    setValues(initialValues)
    setErrors({})
    setIsSubmitting(false)
  }, [initialValues])

  return {
    values,
    errors,
    isSubmitting,
    handleChange,
    handleSubmit,
    setFieldValue,
    setFieldError,
    resetForm,
  }
}
