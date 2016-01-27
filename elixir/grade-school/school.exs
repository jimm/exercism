defmodule School do
  @moduledoc """
  Simulate students in a school.

  Each student is in a grade.
  """

  @doc """
  Add a student to a particular grade in school.
  """
  @spec add(Map.t, String.t, pos_integer) :: Map.t
  def add(db, name, grade) do
    Map.put(db, grade, [name | Map.get(db, grade, [])])
  end

  @doc """
  Return the names of the students in a particular grade.
  """
  @spec grade(Map.t, pos_integer) :: [String]
  def grade(db, grade) do
    Map.get(db, grade, [])
  end

  @doc """
  Sorts the school by grade and name.
  """
  @spec sort(Map) :: Map.t
  def sort(db) do
    db
    |> Enum.map(fn({k, v}) -> {k, Enum.sort(v)} end)
    |> Enum.into(%{})
  end
end
