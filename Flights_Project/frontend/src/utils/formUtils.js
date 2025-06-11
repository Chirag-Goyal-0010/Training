import { useState, useCallback } from 'react';

export const useFormState = (initialState = {}) => {
  const [formData, setFormData] = useState(initialState);
  const [errors, setErrors] = useState({});
  const [touched, setTouched] = useState({});

  const handleChange = useCallback((e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
    // Clear error when field is modified
    setErrors((prev) => ({
      ...prev,
      [name]: '',
    }));
  }, []);

  const handleBlur = useCallback((e) => {
    const { name } = e.target;
    setTouched((prev) => ({
      ...prev,
      [name]: true,
    }));
  }, []);

  const setFieldValue = useCallback((name, value) => {
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
    // Clear error when field is modified
    setErrors((prev) => ({
      ...prev,
      [name]: '',
    }));
  }, []);

  const setFieldError = useCallback((name, error) => {
    setErrors((prev) => ({
      ...prev,
      [name]: error,
    }));
  }, []);

  const setFieldTouched = useCallback((name, isTouched = true) => {
    setTouched((prev) => ({
      ...prev,
      [name]: isTouched,
    }));
  }, []);

  const resetForm = useCallback(() => {
    setFormData(initialState);
    setErrors({});
    setTouched({});
  }, [initialState]);

  const validateForm = useCallback((validationSchema) => {
    const newErrors = {};
    Object.keys(validationSchema).forEach((field) => {
      const value = formData[field];
      const validator = validationSchema[field];
      const error = validator(value);
      if (error) {
        newErrors[field] = error;
      }
    });
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  }, [formData]);

  const isFieldValid = useCallback((name) => {
    return !errors[name] || !touched[name];
  }, [errors, touched]);

  const getFieldError = useCallback((name) => {
    return touched[name] ? errors[name] : '';
  }, [errors, touched]);

  return {
    formData,
    errors,
    touched,
    handleChange,
    handleBlur,
    setFieldValue,
    setFieldError,
    setFieldTouched,
    resetForm,
    validateForm,
    isFieldValid,
    getFieldError,
  };
}; 